<template>
  <div class="">
    <div class="col-12 justify-content-center d-flex flex-row pt-5">
      <div id="div_class" class="flex-item border">
        <h2 class="pt-4 pl-4">Masukkan Alamat</h2>
        <form v-if="hasAddress" @submit.prevent="addAddress" id="addAdress" class="pt-4 pl-4 pr-4">
          <div class="form-group">
            <label>Nama Penerima</label>
            <input type="text" name="penerima" id="penerima" class="form-control" v-model="penerima" required />
          </div>
          <div class="form-group">
            <label>Kabupaten</label>
            <select class="form-control" name="kabupaten" id="kabupaten" v-model="kabupaten" required>
              <option value="Banda Aceh">Banda Aceh</option>
            </select>
          </div>
          <div class="form-row">
            <div class="col">
              <div class="form-group">
                <label>Kecamatan</label>
                <select name="kecamatan" id="kecamatan" class="form-control" v-model="kecamatan" @change="populateDesa"
                  required>
                  <option v-for="option in options" :value="option.value">
                    {{ option.text }}
                  </option>
                </select>
              </div>
            </div>
            <div class="col">
              <div class="form-group">
                <label>Kelurahan</label>
                <select name="kelurahan" id="kelurahan" class="form-control" v-model="kelurahan" required>
                  <option v-for="desa in desaOptions" :value="desa.value">
                    {{ desa.text }}
                  </option>
                </select>
              </div>
            </div>
          </div>
          <div class="form-group">
            <label>Alamat</label>
            <input type="text" name="alamat" id="alamat" class="form-control" v-model="alamat" required />
          </div>
          <div class="form-group">
            <label>No HP</label>
            <input type="text" name="nohp" id="nohp" class="form-control" v-model="nohp" required pattern="\d+"
              @input="validateNumberInput($event, 'nohp')" />
          </div>
          <div class="form-group">
            <label>Catatan</label>
            <textarea name="note" class="form-control" id="note" v-model="note" cols="30" rows="4"></textarea>
          </div>
          <div class="form-group">
            <input type="text" name="id" id="id" value="" class="form-control" hidden />
          </div>
          <div class="form-group">
            <input type="text" name="userId" id="userId" value="" class="form-control" hidden />
          </div>
          <div class="form-row">
            <div class="col">
              <p class="text-center">
                <router-link class="btn btn-dark text-center mx-auto px-5 py-1 mb-2" :to="{ name: 'Cart' }">Kembali ke
                  Keranjang</router-link>
              </p>
            </div>
            <div class="col">
              <p class="text-center">
                <button class="btn btn-dark text-center mx-auto px-5 py-1 mb-2" type="submit">Simpan</button>
              </p>
            </div>
          </div>
        </form>
        <form v-else @submit.prevent="updateAddress" id="updateAdress" class="pt-4 pl-4 pr-4">
          <div class="form-group">
            <label>Nama Penerima</label>
            <input type="text" name="penerima" id="penerima" class="form-control" v-model="penerima" required />
          </div>
          <div class="form-group">
            <label>Kabupaten</label>
            <select name="kabupaten" id="kabupaten" class="form-control" v-model="kabupaten" required>
              <option value="Banda Aceh">Banda Aceh</option>
            </select>
          </div>
          <div class="form-row">
            <div class="col">
              <div class="form-group">
                <label>Kecamatan</label>
                <select name="kecamatan" id="kecamatan" class="form-control" v-model="kecamatan" @change="populateDesa"
                  required>
                  <option v-for="option in options" :value="option.value">
                    {{ option.text }}
                  </option>
                </select>
              </div>
            </div>
            <div class="col">
              <div class="form-group">
                <label>Kelurahan</label>
                <select name="kelurahan" id="kelurahan" class="form-control" v-model="kelurahan" required>
                  <option v-for="desa in desaOptions" :value="desa.value">
                    {{ desa.text }}
                  </option>
                </select>
              </div>
            </div>
          </div>
          <div class="form-group">
            <label>Alamat</label>
            <input type="text" name="alamat" id="alamat" class="form-control" v-model="alamat" required />
          </div>
          <div class="form-group">
            <label>No HP</label>
            <input type="text" name="nohp" id="nohp" class="form-control" v-model="nohp" required pattern="\d+"
              @input="validateNumberInput($event, 'nohp')" />
          </div>
          <div class="form-group">
            <label>Catatan</label>
            <textarea name="note" class="form-control" id="note" v-model="note" cols="30" rows="4"></textarea>
          </div>
          <div class="form-group">
            <input type="text" name="id" id="id" value="" class="form-control" hidden />
          </div>
          <div class="form-group">
            <input type="text" name="userId" id="userId" value="" class="form-control" hidden />
          </div>
          <div class="form-row">
            <div class="col">
              <p class="text-center">
                <router-link class="btn btn-dark text-center mx-auto px-5 py-1 mb-2" :to="{ name: 'Cart' }">Kembali ke
                  Keranjang</router-link>
              </p>
            </div>
            <div class="col">
              <p class="text-center">
                <button class="btn btn-dark text-center mx-auto px-5 py-1 mb-2" type="submit">Simpan</button>
              </p>
            </div>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      id: null,
      penerima: null,
      kabupaten: null,
      kecamatan: null,
      kelurahan: null,
      alamat: null,
      nohp: null,
      note: null,
      token: localStorage.getItem('token'),
      address: [],
      options: [
        { text: 'Baiturrahman', value: 'Baiturrahman' },
        { text: 'Banda Raya', value: 'Banda Raya' },
        { text: 'Jaya Baru', value: 'Jaya Baru' },
        { text: 'Kuta Alam', value: 'Kuta Alam' },
        { text: 'Kuta Raja', value: 'Kuta Raja' },
        { text: 'Lueng Bata', value: 'Lueng Bata' },
        { text: 'Meuraxa', value: 'Meuraxa' },
        { text: 'Syiah Kuala', value: 'Syiah Kuala' },
        { text: 'Ulee Kareng', value: 'Ulee Kareng' }
      ],
      desaOptions: []
    };
  },
  props: ["baseURL", "address"],
  computed: {
    hasAddress() {
      return this.id === null;
    }
  },
  methods: {
    populateDesa() {
      const kecamatan = this.kecamatan;
      const desaOptions = [];

      switch (kecamatan) {
        case "Banda Raya":
          desaOptions.push({ text: "Lhong Raya", value: "Lhong Raya" });
          desaOptions.push({ text: "Lampeout", value: "Lampeout" });
          desaOptions.push({ text: "Mibo", value: "Mibo" });
          desaOptions.push({ text: "Lam Ara", value: "Lam Ara" });
          desaOptions.push({ text: "Lhong Cut", value: "Lhong Cut" });
          desaOptions.push({ text: "Peunyerat", value: "Peunyerat" });
          desaOptions.push({ text: "Geuceu Kayee Jato", value: "Geuceu Kayee Jato" });
          desaOptions.push({ text: "Geucu Iniem", value: "Geucu Iniem" });
          desaOptions.push({ text: "Lamlagang", value: "Lamlagang" });
          break;
        case "Meuraxa":
          desaOptions.push({ text: "Ulee Lheu", value: "Ulee Lheu" });
          desaOptions.push({ text: "Deah Geulumpang", value: "Deah Geulumpang" });
          desaOptions.push({ text: "Gampong Pie", value: "Gampong Pie" });
          desaOptions.push({ text: "Lambung", value: "Lambung" });
          desaOptions.push({ text: "Alue Deah Teungoh", value: "Alue Deah Teungoh" });
          desaOptions.push({ text: "Deah Baro", value: "Deah Baro" });
          desaOptions.push({ text: "Cot Lamkeuweuh", value: "Cot Lamkeuweuh" });
          desaOptions.push({ text: "Blang Oi", value: "Blang Oi" });
          desaOptions.push({ text: "Gampong Blang", value: "Gampong Blang" });
          desaOptions.push({ text: "Lamjabat", value: "Lamjabat" });
          desaOptions.push({ text: "Asoe Nanggroe", value: "Asoe Nanggroe" });
          desaOptions.push({ text: "Punge Ujong", value: "Punge Ujong" });
          desaOptions.push({ text: "Surien", value: "Surien" });
          desaOptions.push({ text: "Gampong Baro", value: "Gampong Baro" });
          desaOptions.push({ text: "Lampaseh Aceh", value: "Lampaseh Aceh" });
          desaOptions.push({ text: "Punge Jurong", value: "Punge Jurong" });
          break;
        case "Ulee Kareng":
          desaOptions.push({ text: "Curieh", value: "Curieh" });
          desaOptions.push({ text: "Lamteh", value: "Lamteh" });
          desaOptions.push({ text: "Lambhuk", value: "Lambhuk" });
          desaOptions.push({ text: "Ilie", value: "Ilie" });
          desaOptions.push({ text: "Pango Deah", value: "Pango Deah" });
          desaOptions.push({ text: "Pango Raya", value: "Pango Raya" });
          desaOptions.push({ text: "Lam Glumpang", value: "Lam Glumpang" });
          desaOptions.push({ text: "Ie Masen Ulee Kareng", value: "Ie Masen Ulee Kareng" });
          desaOptions.push({ text: "Doy", value: "Doy" });
          break;
        case "Kuta Alam":
          desaOptions.push({ text: "Bandar Baru", value: "Bandar Baru" });
          desaOptions.push({ text: "Kota Baru", value: "Kota Baru" });
          desaOptions.push({ text: "Beurawe", value: "Beurawe" });
          desaOptions.push({ text: "Kuta Alam", value: "Kuta Alam" });
          desaOptions.push({ text: "Keuramat", value: "Keuramat" });
          desaOptions.push({ text: "Laksana", value: "Laksana" });
          desaOptions.push({ text: "Mulia", value: "Mulia" });
          desaOptions.push({ text: "Peunayong", value: "Peunayong" });
          desaOptions.push({ text: "Lampulo", value: "Lampulo" });
          desaOptions.push({ text: "Lamdingin", value: "Lamdingin" });
          desaOptions.push({ text: "Lambaro Skep", value: "Lambaro Skep" });
          break;
        case "Syiah Kuala":
          desaOptions.push({ text: "Lamgugop", value: "Lamgugop" });
          desaOptions.push({ text: "Jeulingke", value: "Jeulingke" });
          desaOptions.push({ text: "Tibang", value: "Tibang" });
          desaOptions.push({ text: "Alue Naga", value: "Alue Naga" });
          desaOptions.push({ text: "Pineung", value: "Pineung" });
          desaOptions.push({ text: "Deyah Raya", value: "Deyah Raya" });
          desaOptions.push({ text: "Ie Masen Kayee Adang", value: "Ie Masen Kayee Adang" });
          desaOptions.push({ text: "Kopelma Darussalam", value: "Kopelma Darussalam" });
          desaOptions.push({ text: "Rukoh", value: "Rukoh" });
          break;
        case "Baiturrahman":
          desaOptions.push({ text: "Neusu Jaya", value: "Neusu Jaya" });
          desaOptions.push({ text: 'Kampung Baru', value: 'Kampung Baru' });
          desaOptions.push({ text: 'Peuniti', value: 'Peuniti' });
          desaOptions.push({ text: 'Seutui', value: 'Seutui' });
          desaOptions.push({ text: 'Ateuk Pahlawan', value: 'Ateuk Pahlawan' });
          desaOptions.push({ text: 'Sukaramai', value: 'Sukaramai' });
          desaOptions.push({ text: 'Neusu Aceh', value: 'Neusu Aceh' });
          desaOptions.push({ text: 'Ateuk Munjeng', value: 'Ateuk Munjeng' });
          desaOptions.push({ text: 'Ateuk Jawo', value: 'Ateuk Jawo' });
          desaOptions.push({ text: 'Ateuk Deah Tanoh', value: 'Ateuk Deah Tanoh' });
          break;
        case "Lueng Bata":
          desaOptions.push({ text: 'Lueng Bata', value: 'Lueng Bata' });
          desaOptions.push({ text: 'Cot Mesjid', value: 'Cot Mesjid' });
          desaOptions.push({ text: 'Pante Riek', value: 'Pante Riek' });
          desaOptions.push({ text: 'Blang Cut', value: 'Blang Cut' });
          desaOptions.push({ text: 'Lamseupeung', value: 'Lamseupeung' });
          desaOptions.push({ text: 'Batoh', value: 'Batoh' });
          desaOptions.push({ text: 'Suka Damai', value: 'Suka Damai' });
          desaOptions.push({ text: 'Lamdom', value: 'Lamdom' });
          desaOptions.push({ text: 'Lampaloh', value: 'Lampaloh' });
          break;
        case "Kuta Raja":
          desaOptions.push({ text: 'Keudah', value: 'Keudah' });
          desaOptions.push({ text: 'Peulanggahan', value: 'Peulanggahan' });
          desaOptions.push({ text: 'Merduati', value: 'Merduati' });
          desaOptions.push({ text: 'Lampaseh Kota', value: 'Lampaseh Kota' });
          desaOptions.push({ text: 'Gampong Jawa', value: 'Gampong Jawa' });
          desaOptions.push({ text: 'Gampong Pande', value: 'Gampong Pande' });
          desaOptions.push({ text: 'Mulia', value: 'Mulia' });
          break;
        case "Jaya Baru":
          desaOptions.push({ text: 'Lampoh Daya', value: 'Lampoh Daya' });
          desaOptions.push({ text: 'Emperom', value: 'Emperom' });
          desaOptions.push({ text: 'Lamjamee', value: 'Lamjamee' });
          desaOptions.push({ text: 'Bitai', value: 'Bitai' });
          desaOptions.push({ text: 'Lamteumen Barat', value: 'Lamteumen Barat' });
          desaOptions.push({ text: 'Lamteumen Timur', value: 'Lamteumen Timur' });
          desaOptions.push({ text: 'Ulee Pata', value: 'Ulee Pata' });
          desaOptions.push({ text: 'Geuceu Meunara', value: 'Geuceu Meunara' });
          desaOptions.push({ text: 'Punge Blang Cut', value: 'Punge Blang Cut' });
          break;
        default:
          desaOptions.push({ text: "Please select Kecamatan first", value: "" });
      }

      this.desaOptions = desaOptions;

      // If kelurahan was already set, ensure it is selected correctly
      if (this.address.kelurahan) {
        this.kelurahan = this.address.kelurahan;
      }
    },
    fetchAddress() {
      axios.get(`${this.baseURL}api/address/users`, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        }
      }).then((response) => {
        if (response.data.code === 200 && response.data.data.length > 0) {
          const result = response.data.data[0];
          this.id = result.id;
          this.penerima = result.nama_penerima;
          this.kabupaten = result.kabupaten;
          this.kecamatan = result.kecamatan;
          this.kelurahan = result.kelurahan;
          this.alamat = result.alamat;
          this.nohp = result.nohp;
          this.note = result.note;
          console.log(result);

          this.$nextTick(() => {
            this.populateDesa();
          });
        } else {
          this.id = null;
        }
      }).catch((err) => {
        console.log(err);
      })
    },
    addAddress() {
      axios.post(`${this.baseURL}api/address`, {
        nama_penerima: this.penerima,
        kabupaten: this.kabupaten,
        kecamatan: this.kecamatan,
        kelurahan: this.kelurahan,
        alamat: this.alamat,
        nohp: this.nohp,
        note: this.note,
      }, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        },
      }).then((response) => {
        if (response.data.code === 200) {
          this.fetchAddress();
          console.log(response.data.data);
          this.$emit("fecthData");
          swal("Success", "Address added successfully", "success");
        }
      }).catch((error) => {
        console.log(error);
      })
    },
    updateAddress() {
      console.log(this.id);
      axios.put(`${this.baseURL}api/address/update`, {
        id: this.id,
        nama_penerima: this.penerima,
        kabupaten: this.kabupaten,
        kecamatan: this.kecamatan,
        kelurahan: this.kelurahan,
        alamat: this.alamat,
        nohp: this.nohp,
        note: this.note,
      }, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        }
      }).then((response) => {
        if (response.data.code === 200) {
          console.log(response.data.data);
          this.fetchAddress();
          swal("Success", "Address updated successfully", "success");
        }
      }).catch((err) => {
        console.log(err);
      });
    },
    validateNumberInput(event, field) {
      const regex = /^\d*$/;
      if (!regex.test(event.target.value)) {
        this[field] = event.target.value.replace(/[^\d]/g, '');
      }
    }, 
  },
  mounted() {
    this.fetchAddress();
    this.token = localStorage.getItem("token");
  }
}
</script>

<style>
@media only screen and (min-width: 992px) {
  #div_class {
    width: 40%;
  }
}
</style>