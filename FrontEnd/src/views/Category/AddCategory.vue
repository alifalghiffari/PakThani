<template>
  <div class="container">
    <div class="row">
      <div class="col-sm-5">
        <div class="d-flex justify-content-between align-items-center mb-3 mt-3">
          <button type="button" class="btn btn-secondary" @click="goBack">Kembali</button>
        </div>
      </div>
      <div class="col text-start">
        <h4 class="pt-3">Tambah Kategori Baru</h4>
      </div>
    </div>

    <div class="row">
      <div class="col-3"></div>
      <div class="col-md-6 px-5 px-md-0">
        <form>
          <div class="form-group">
            <label>Nama kategori</label>
            <input type="text" class="form-control" v-model="categoryName" @input="validateCharacterInput" required>
          </div>
          <div class="form-group">
            <label for="imageURL">Gambar</label>
            <input type="file" class="form-control-file" id="imageURL" ref="imageURL" @change="handleFileUploadMain">
          </div>
          <button type="button" class="btn btn-primary" @click="addCategory">Simpan</button>
        </form>
      </div>
      <div class="col-3"></div>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      categoryName: null,
      imageURL: null,
      imgMain: '',
      token: localStorage.getItem('token'),
    }
  },
  props: ["baseURL", "categories"],
  methods: {
    addCategory() {
      const slug = this.generateSlug(this.categoryName);

      if (this.$refs.imageURL.files.length > 0) {
        const file = this.$refs.imageURL.files[0];
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => {
          let base64 = reader.result.split(',')[1];
          this.imgMain = base64;
        };
      }

      const newCategory = {
        category: this.categoryName,
        image: this.imgMain,
        slug: slug,
      }

      axios.post(`${this.baseURL}api/categories`, newCategory, {
        headers: {
          Authorization: `Bearer ${this.token}`,
        },
      }).then((res) => {
        if (res.data.code === 200) {
          swal({
            text: "Produk berhasil ditambahkan!",
            icon: "success",
            closeOnClickOutside: false,
          });
          this.$emit("fetchData");
          this.$router.push({ name: 'AdminCategory' });
        }
      }).catch(err => {
        console.error(err);
      });
    },
    generateSlug(name) {
      // Remove special characters and replace spaces with hyphens
      const slug = name.toLowerCase().replace(/[^a-z0-9 ]/g, '').replace(/ /g, '-');
      return slug;
    },
    handleFileUploadMain(event) {
      this.imageURL = event.target.files;
    },
    validateCharacterInput(event) {
      const regex = /^[a-zA-Z\s]*$/;
      if (!regex.test(event.target.value)) {
        this.categoryName = event.target.value.replace(/[^a-zA-Z\s]/g, '');
      }
    },
    goBack() {
      this.$router.go(-1);
    }
  },
  mounted() {
    if (!this.token) {
      this.$router.push({ name: 'Signin' });
    }
  }
}
</script>

<style scoped>
h4 {
  font-family: 'Roboto', sans-serif;
  color: #484848;
  font-weight: 700;
}
</style>
