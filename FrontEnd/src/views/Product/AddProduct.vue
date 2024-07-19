<script setup>
import Editor from '@tinymce/tinymce-vue'
</script>

<template>
  <div class="container">
    <div class="row">
      <div class="col-sm-5">
        <div class="d-flex justify-content-between align-items-center mb-3 mt-3">
          <button type="button" class="btn btn-secondary" @click="goBack">Kembali</button>
        </div>
      </div>
      <div class="col text-start">
        <h4 class="pt-3">Tambah Produk Baru</h4>
      </div>
    </div>

    <div class="row">
      <div class="col-3"></div>
      <div class="col-md-6 px-5 px-md-0">
        <form @submit.prevent="addProduct">
          <div class="form-group">
            <label>Kategori</label>
            <select class="form-control" v-model.number="categoryId" required>
              <option v-for="category in categories" :value="category.id">
                {{ category.category }}
              </option>
            </select>
          </div>
          <div class="form-group">
            <label>Nama</label>
            <input type="text" class="form-control" v-model="name" @input="validateCharacterInput" required>
          </div>
          <div class="form-group">
            <label>Deskripsi</label>
            <Editor api-key="jzlaylf6k5nu60l768s6m459rg56swoeupz40pf401h1gi6o" v-model="description" :init="{
              toolbar_mode: 'sliding',
              plugins: 'anchor autolink charmap codesample emoticons image link lists media searchreplace table visualblocks wordcount linkchecker',
              toolbar: 'undo redo | blocks fontfamily fontsize | bold italic underline strikethrough | link image media table | align lineheight | numlist bullist indent outdent | emoticons charmap | removeformat',
            }" />
          </div>
          <div class="form-group">
            <label for="imageURL">Gambar Utama</label>
            <input type="file" class="form-control-file" id="imageURL" ref="imageURL" @change="handleFileUploadMain">
          </div>
          <div class="custom-file mb-3">
            <input multiple type="file" class="custom-file-input" id="image1" name="image1" ref="image1"
              @change="handleFileUpload">
            <label class="custom-file-label" for="image1">Opsi Gambar</label>
          </div>
          <div class="form-group">
            <label>Harga</label>
            <input type="number" class="form-control" v-model.number="price" required pattern="\d+"
              @input="validateNumberInput($event, 'price')">
          </div>
          <div class="form-group">
            <label>ketersediaan</label>
            <input type="number" class="form-control" v-model.number="stock" required pattern="\d+"
              @input="validateNumberInput($event, 'stock')">
          </div>
          <button type="submit" class="btn btn-primary">Simpan</button>
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
      categoryId: null,
      name: null,
      description: null,
      imageURL: null,
      price: null,
      stock: null,
      image1: [],
      categories: [],
      imgMain: '',
      token: localStorage.getItem('token'),
    }
  },
  props: ["baseURL"],
  methods: {
    getCategories() {
      axios.get(`${this.baseURL}api/category`, {
        headers: {
          Authorization: `Bearer ${this.token}`
        }
      }).then((res) => {
        if (res.data.code === 200) {
          this.categories = res.data.data;
        }
      }).catch(err => {
        console.error(err);
      });
    },
    addProduct() {
      const slug = this.generateSlug(this.name);
      const base64Images = [];

      if (this.$refs.imageURL.files.length > 0) {
        const file = this.$refs.imageURL.files[0];
        const reader = new FileReader();
        reader.readAsDataURL(file);
        reader.onload = () => {
          let base64 = reader.result.split(',')[1];
          this.imgMain = base64;

          this.handleOptionalImages(slug, base64Images);
        };
      } else {
        this.handleOptionalImages(slug, base64Images);
      }
    },
    handleOptionalImages(slug, base64Images) {
      const files = this.$refs.image1.files;

      if (files.length > 0) {
        for (let i = 0; i < files.length; i++) {
          const file = files[i];
          const reader = new FileReader();
          reader.readAsDataURL(file);
          reader.onload = () => {
            let base64String = reader.result.split(',')[1];
            base64Images.push(base64String);

            if (base64Images.length === files.length) {
              this.submitProduct(slug, base64Images);
            }
          };
        }
      } else {
        this.submitProduct(slug, base64Images);
      }
    },
    submitProduct(slug, base64Images) {
      const newProduct = {
        category_id: this.categoryId,
        name: this.name,
        description: this.description,
        image: this.imgMain,
        price: this.price,
        quantity: this.stock,
        slug: slug,
        images: base64Images,
      };

      axios.post(`${this.baseURL}api/products`, newProduct, {
        headers: {
          Authorization: `Bearer ${this.token}`,
          'Content-Type': 'application/json',
        },
      }).then((res) => {
        if (res.data.code === 200) {
          swal({
            text: "Produk berhasil ditambahkan!",
            icon: "success",
            closeOnClickOutside: false,
          });
          this.$emit("fetchData");
          this.$router.push({ name: 'AdminProduct' });
        }
      }).catch(err => {
        console.error(err);
      });
    },
    generateSlug(name) {
      const slug = name.toLowerCase().replace(/[^a-z0-9 ]/g, '').replace(/ /g, '-');
      return slug;
    },
    handleFileUploadMain(event) {
      this.imageURL = event.target.files;
    },
    handleFileUpload(event) {
      this.image1 = event.target.files;
    },
    validateNumberInput(event, field) {
      const regex = /^\d*$/;
      if (!regex.test(event.target.value)) {
        this[field] = event.target.value.replace(/[^\d]/g, '');
      }
    },
    validateCharacterInput(event) {
      const regex = /^[a-zA-Z\s]*$/;
      if (!regex.test(event.target.value)) {
        this.name = event.target.value.replace(/[^a-zA-Z\s]/g, '');
      }
    },
    goBack() {
      this.$router.go(-1);
    }
  },
  mounted() {
    if (!this.token) {
      this.$router.push({ name: 'Signin' });
    } else {
      this.getCategories();
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
