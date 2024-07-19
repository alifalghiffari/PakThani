<template>
  <div class="container">
    <div class="row">
      <div class="col-12 text-center">
        <h4 class="pt-3">{{category.categoryName}}</h4>
        <h5></h5>
      </div>
    </div>

    <div class="row">
      <!-- <img v-show="len == 0" class="img-fluid" src="../../assets/sorry.jpg" alt="Sorry"> -->
      <div v-for="product in produk" :key="product.id" class="col-md-6 col-xl-4 col-12 pt-3  justify-content-around d-flex">
        <ProductBox :product="product">
        </ProductBox> 
      </div>
      <p></p>
    </div>
  </div>
</template>

<script>
import ProductBox from '../../components/Product/ProductBox';

export default {
  name: 'ListProducts',
  data() {
    return {
      id: null,
      categories: [], // Mengubah inisialisasi categories menjadi array kosong
      len: 0,
      produk : [],
      msg: null
    }
  },
  components: { ProductBox },
  props: ["baseURL", "category", "products"],
  mounted() {
    this.id = this.$route.params.id;
    console.log(this.id);
    // Filter kategori yang sesuai dengan ID
    this.categories = this.category.filter((category) => category.id == this.id);
    this.produk = this.products.filter((product) => product.categoryId == this.id);
    console.log(this.produk);



    if (this.categories.length === 0) {
      this.msg = "Sorry, no categories found for this ID";
      return;
    }

    // Menghitung total produk dari seluruh kategori yang sesuai dengan ID
    this.len = this.categories.reduce((acc, category) => {
      return acc + (category.products ? category.products.length : 0);
    }, 0);
    
    if (this.len === 0) {
      this.msg = "Sorry, no products found for these categories";
    } else if (this.len === 1) {
      this.msg = "Only 1 product found";
    } else {
      this.msg = this.len + " products found";
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

h5 {
  font-family: 'Roboto', sans-serif;
  color: #686868;
  font-weight: 300;
}
</style>
