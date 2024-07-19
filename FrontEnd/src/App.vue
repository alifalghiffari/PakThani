<template>
  <Navbar :cartCount="cartCount" @resetCartCount="resetCartCount"
    v-if="!['Signup', 'Signin', 'AboutUs', 'ContactUs', 'OurBusiness', 'VerifyAccount', 'Verify', 'VerifyOTP', 'ResetPassword'].includes($route.name)" />
  <div style="min-height: 60vh">
    <router-view v-if="products && category" :baseURL="baseURL" :products="products" :category="category"
      @fetchData="fetchData">
    </router-view>
  </div>
  <Footer v-if="!['Signup', 'Signin', 'ShowDetails', 'VerifyAccount', 'Verify', 'VerifyOTP', 'ResetPassword'].includes($route.name)" />
</template>

<script>
import Navbar from './components/Navbar.vue';
import Footer from './components/Footer.vue';
export default {
  data() {
    return {
      // baseURL: 'https://limitless-lake-55070.herokuapp.com/',
      baseURL: "http://localhost:3000/",
      products: null,
      category: null,
      key: 0,
      token: null,
      cartCount: 0,
    };
  },

  components: { Footer, Navbar },
  methods: {
    async fetchData() {
      // fetch products
      await axios
        .get(this.baseURL + '/api/products')
        .then((res) => (this.products = res.data.data))
        .catch((err) => console.log(err));

      //fetch categories
      await axios
        .get(this.baseURL + '/api/categories')
        .then((res) => (this.category = res.data.data))
        .catch((err) => console.log(err));

      //fetch cart items
      // if (this.token) {
      //   await axios.get(`${this.baseURL}cart/?token=${this.token}`).then(
      //     (response) => {
      //       if (response.status == 200) {
      //         // update cart
      //         this.cartCount = Object.keys(response.data.cartItems).length;
      //       }
      //     },
      //     (error) => {
      //       console.log(error);
      //     }
      //   );
      // }
    },
    resetCartCount() {
      this.cartCount = 0;
    },
  },
  mounted() {
    this.token = localStorage.getItem('token');
    this.fetchData();
  },
};
</script>

<style>
html {
  overflow-y: scroll;
}
</style>
