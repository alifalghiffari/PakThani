<template>
  <div class="container">
    <div class="row">
      <div class="col-12 text-center">
        <h4 class="pt-3">Pesanan Masuk</h4>
      </div>
    </div>
    <!--        for each order display -->
    <div v-for="order in orderList" :key="order.id" class="row mt-2 pt-3 justify-content-around">
      <!-- <div class="col-2"></div> -->
      <!-- <div class="col-md-3 embed-responsive embed-responsive-16by9"> -->
      <!--                display image in left-->
      <!-- <img v-bind:src="order.imageURL" class="w-100 card-img-top embed-responsive-item"> -->
      <!-- </div> -->
      <div class="col-md-5 px-3">
        <div class="card-block px-3">
          <h6 class="card-title">
            <router-link @click.native v-bind:to="'/admin/order/' + order.id">Nomor Pesanan : {{ order.id
              }}</router-link>
          </h6>
          <p class="mb-0">{{ order.total_items }} Item<span v-if="order.total_items > 1"></span></p>
          <p id="item-price" class="mb-0 font-weight-bold">Jumlah Biaya : Rp {{ order.total_price }}</p>
          <!-- <p id="item-username">Ordered By : {{order.username}}</p> -->
          <form @submit.prevent="submitStatus(order)">
            <div class="row align-items-end">
              <div class="col-8">
                <label for="status" class="mb-0">Status : {{ order.order_status }}</label>
                <select name="status" id="status" class="form-control" v-model="order.selectedStatus" required>
                  <option disabled>Pilih Status</option>
                  <option value="PROCESS">Dalam Pengiriman</option>
                  <option value="DELIVERED">Selesai</option>
                </select>
              </div>
              <div class=" col">
                <button type="submit" class="btn btn-success">Simpan</button>
              </div>
            </div>
          </form>
          <!-- <p id="item-status" class="mb-0 ">Status : {{ order.order_status }}</p> -->
          <p id="item-status">Catatan : {{ order.address.note }}</p>
          <a href="" @click.prevent="showAddressModal(order)">Alamat</a>
        </div>
      </div>
      <div class="col-2"></div>
      <div class="col-12">
        <hr>
      </div>
    </div>
  </div>

  <!-- Modal for Detail Address -->
  <div class="modal fade" id="addressModal" tabindex="-1" role="dialog" aria-labelledby="addressModalLabel"
    aria-hidden="true">
    <div class="modal-dialog" role="document">
      <div class="modal-content">
        <div class="modal-header">
          <h5 class="modal-title" id="addressModalLabel">Alamat Lengkap</h5>
          <button type="button" class="close" data-dismiss="modal" aria-label="Close">
            <span aria-hidden="true">&times;</span>
          </button>
        </div>
        <div class="modal-body">
          <p>Nama : {{ selectedAddress.nama_penerima }}</p>
          <p>Kabupaten : {{ selectedAddress.kabupaten }}</p>
          <p>Alamat : {{ selectedAddress.alamat }}, {{ selectedAddress.kelurahan }} {{ selectedAddress.kecamatan }}</p>
          <p>Note : {{ selectedAddress.note }}</p>
          <p>No HP : {{ selectedAddress.nohp }}</p>
        </div>
        <div class="modal-footer">
          <button type="button" class="btn btn-secondary" data-dismiss="modal">Tutup</button>
        </div>
      </div>
    </div>
  </div>

</template>

<script>
const axios = require('axios')
export default {

  data() {
    return {
      token: null,
      orderList: [],
      status: 'Pilih Status',
      selectedAddress: {}
    }
  },
  props: ["baseURL"],
  name: 'Order',
  methods: {
    fecthOrder() {
      axios.get(`${this.baseURL}api/order`, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        },
      }).then((response) => {
        if (response.status === 200) {
          this.orderList = response.data.data.map(order => ({
            ...order,
            selectedStatus: order.order_status
          }));
        }
      }).catch((err) => {
        console.error(err);
      });
    },
    showAddressModal(order) {
      this.selectedAddress = order.address;
      $('#addressModal').modal('show');
    },
    submitStatus(order) {
      const update = {
        id: order.id,
        order_status: order.selectedStatus,
        payment_status: order.selectedStatus === 'DELIVERED' ? 'SUCCESS' : 'PENDING',
      };

      axios.put(`${this.baseURL}api/orders/${order.id}`, update, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        },
      }).then((res) => {
        if (res.data.code === 200) {
          swal({
            text: "Status berhasil diubah!",
            icon: "success",
            closeOnClickOutside: false,
          });
        }
      }).catch((err) => {
        console.error(err);
      });
    },
  },

  mounted() {
    this.token = localStorage.getItem("token");
    this.fecthOrder();
  },
};

</script>

<style scoped>
h4,
h5 {
  font-family: 'Roboto', sans-serif;
  color: #484848;
  font-weight: 700;
}

.embed-responsive .card-img-top {
  object-fit: cover;
}
</style>