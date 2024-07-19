<template>
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
    <router-link class="navbar-brand" @click.native :to="{ name: 'Home' }">
      <img id="logo" src="../assets/logo.png" />
    </router-link>

    <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent"
      aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>

    <div class="collapse navbar-collapse" id="navbarSupportedContent">
      <form class="form-inline ml-auto mr-auto" ref="searchForm">
        <div class="input-group">
          <input size="95" type="text" class="form-control" placeholder="Cari Produk ..." aria-label="Username"
            aria-describedby="basic-addon1" v-model="searchQuery" @keyup="searchProduct(searchQuery)"
            @input="updateDropdownWidth" />
          <div class="input-group-prepend">
            <span class="input-group-text" id="search-button-navbar">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-search"
                viewBox="0 0 16 16">
                <path
                  d="M11.742 10.344a6.5 6.5 0 1 0-1.397 1.398h-.001c.03.04.062.078.098.115l3.85 3.85a1 1 0 0 0 1.415-1.414l-3.85-3.85a1.007 1.007 0 0 0-.115-.1zM12 6.5a5.5 5.5 0 1 1-11 0 5.5 5.5 0 0 1 11 0z" />
              </svg>
            </span>
          </div>
        </div>

        <div v-if="filteredProducts && filteredProducts.length > 0" class="search-results"
          :style="{ width: dropdownWidth + 'px' }">
          <ul class="list-group">
            <li class="list-group-item" v-for="item in filteredProducts" :key="item.id">
              <div class="container">
                <div class="row row-cols-3">
                  <div class="col-2">
                    <img :src="getImagePathMain(item.image)" alt="" />
                  </div>
                  <div class="col-sm-7">
                    <p class="mb-0">{{ item.name }}</p>
                    <div class="row">
                      <div class="col-8 col-sm-6">
                        <p>Rp {{ item.price }}</p>
                      </div>
                    </div>
                  </div>
                  <div class="col-sm text-right">
                    <router-link class="text-warning" :to="{ name: 'ShowDetails', params: { id: item.id } }"
                      @click.native="clearSearch">Tampilkan Detail</router-link>
                  </div>
                </div>
              </div>
            </li>
          </ul>
        </div>
      </form>

      <!-- DropDowns -->
      <ul class="navbar-nav ml-auto">
        <li class="nav-item dropdown">
          <a class="nav-link text-light dropdown-toggle" href="#" id="navbarDropdown" role="button"
            data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
            Jelajahi
          </a>
          <div class="dropdown-menu" aria-labelledby="navbarDropdown">
            <router-link class="dropdown-item" @click.native :to="{ name: 'Home' }">Home</router-link>
            <router-link class="dropdown-item" @click.native :to="{ name: 'Product' }">Produk</router-link>
            <router-link class="dropdown-item" @click.native :to="{ name: 'Category' }">Kategori</router-link>
          </div>
        </li>

        <li class="nav-item dropdown">
          <a class="nav-link text-light dropdown-toggle" href="#" id="navbarDropdown" role="button"
            data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
            Akun
          </a>
          <div class="dropdown-menu" aria-labelledby="navbarDropdown">
            <router-link class="dropdown-item" v-if="!token" @click.native :to="{ name: 'Signin' }">Masuk</router-link>
            <router-link class="dropdown-item" v-if="!token" @click.native :to="{ name: 'Signup' }">Daftar</router-link>
            <a class="dropdown-item" v-if="token" href="#" @click="signout">Keluar</a>
          </div>
        </li>

        <li class="nav-item">
          <router-link class="nav-link text-light" v-if="token && !isAdmin" @click.native
            :to="{ name: 'Order' }">Transaksi</router-link>
        </li>
        <li class="nav-item">
          <div class="">
            <router-link class="nav-link text-light pl-3 pr-3" v-if="token && isAdmin"
              :to="{ name: 'Admin' }">Admin</router-link>
          </div>
          <div id="cart">
            <router-link class="text-light" v-if="token && !isAdmin" @click.native :to="{ name: 'Cart' }"><i
                class="fa fa-shopping-cart" style="font-size:36px"></i></router-link>
          </div>
        </li>
      </ul>
    </div>
  </nav>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Navbar',
  props: ['cartCount'],
  data() {
    return {
      token: null,
      products: [],
      filteredProducts: [],
      searchQuery: '',
      isLoading: false,
      dropdownWidth: 0, // New data property
    };
  },
  methods: {
    signout() {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
      this.token = null;
      this.$emit('resetCartCount');
      this.$router.push({ name: 'Home' });
      swal({
        text: 'Anda Telah Keluar. Silahkan Datang Kembali!',
        icon: 'success',
        closeOnClickOutside: false,
      });
    },
    getProduct() {
      this.isLoading = true;
      axios
        .get(`http://localhost:3000/api/products`)
        .then((res) => {
          if (res.data.code === 200) {
            this.products = res.data.data;
          }
        })
        .catch((err) => {
          console.error(err);
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
    searchProduct(query) {
      this.searchQuery = query;
      if (query.trim() === '') {
        this.filteredProducts = [];
      } else {
        if (this.products && this.products.length > 0) {
          this.filteredProducts = this.products.filter((product) =>
            product.name.toLowerCase().includes(query.toLowerCase())
          );
        } else {
          this.filteredProducts = [];
        }
      }
    },
    getImagePathMain(image) {
      try {
        return require(`../assets/img-main/${image}`);
      } catch (e) {
        return ''; 
      }
    },
    updateDropdownWidth() {
      this.$nextTick(() => {
        const formWidth = this.$refs.searchForm.clientWidth;
        this.dropdownWidth = formWidth;
      });
    },
    clearSearch() {
      this.searchQuery = '';
      this.filteredProducts = [];
    },
  },
  computed: {
    isAdmin() {
      const role = localStorage.getItem('user');
      return role === 'admin';
    },
  },
  mounted() {
    this.getProduct();
    this.token = localStorage.getItem('token');
    this.updateDropdownWidth();
    window.addEventListener('resize', this.updateDropdownWidth);
  },
  beforeDestroy() {
    window.removeEventListener('resize', this.updateDropdownWidth);
  },
};
</script>

<style scoped>
#logo {
  width: 150px;
  margin-left: 20px;
  margin-right: 20px;
}

.nav-link {
  color: rgba(255, 255, 255);
}

#search-button-navbar {
  background-color: #febd69;
  border-color: #febd69;
  border-top-right-radius: 2px;
  border-bottom-right-radius: 2px;
}

#nav-cart-count {
  background-color: red;
  color: white;
  border-radius: 50%;
  position: absolute;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 15px;
  height: 15px;
  font-size: 15px;
  margin-left: 10px;
}

#cart {
  position: relative;
}

.search-results {
  top: 50px;
  margin-top: 6px;
  position: absolute;
  background-color: #ffffff; /* Changed to white for better contrast */
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
  max-height: 300px;
  overflow-y: auto;
  z-index: 2;
  width: 100%; /* Relative width */
}

.search-results .list-group-item {
  border: none;
}

.search-results .list-group-item img {
  width: 90px;
  height: 50px; 
  object-fit: cover; 
}

.search-results .list-group-item a {
  display: block;
  padding: 8px 16px;
  color: #333;
  text-decoration: none;
}

.search-results .list-group-item a:hover {
  background-color: #f5f5f5;
}
</style>
