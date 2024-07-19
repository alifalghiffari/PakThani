<template>
  <div id="home">
    <!-- Page Wrapper -->
    <div id="carouselBackground" class="carousel slide page-holder bg-cover" data-ride="carousel">
      <div class="carousel-inner">
        <div class="carousel-item active">
          <div class="page-holder bg-cover" id="bg1">
            <div class="container py-5">
              <header class="text-left text-white py-5">
                <!-- Slide content here -->
              </header>
            </div>
          </div>
        </div>
        <div class="carousel-item">
          <div class="page-holder bg-cover" id="bg2">
            <div class="container py-5">
              <header class="text-left text-white py-5">
                <!-- Slide content here -->
              </header>
            </div>
          </div>
        </div>
        <div class="carousel-item">
          <div class="page-holder bg-cover" id="bg3">
            <div class="container py-5">
              <header class="text-left text-white py-5">
                <!-- Slide content here -->
              </header>
            </div>
          </div>
        </div>
      </div>
      <a class="carousel-control-prev" href="#carouselBackground" role="button" data-slide="prev">
        <span class="carousel-control-prev-icon" aria-hidden="true"></span>
        <span class="sr-only">Previous</span>
      </a>
      <a class="carousel-control-next" href="#carouselBackground" role="button" data-slide="next">
        <span class="carousel-control-next-icon" aria-hidden="true"></span>
        <span class="sr-only">Next</span>
      </a>
      <div class="fixed-button-container">
        <a href="#start-shopping" role="button" class="bg-warning px-2 py-2 p-4 fixed-button" id="heading">Mulai
          Belanja</a>
      </div>
    </div>

    <div id="start-shopping" class="container">
      <div class="row">
        <div class="col-12 text-left">
          <h2 class="pt-3">Kategori</h2>
        </div>
      </div>
      <div class="row">
        <div v-for="index in this.category_size" :key="index"
          class="col-md-6 col-xl col-12 pt-3  justify-content-around d-flex">
          <CategoryBox :category="category[index - 1]">
          </CategoryBox>
        </div>
      </div>
    </div>

    <hr>
    <div class="container">
      <div class="row">
        <div class="col-12 text-left">
          <h2 class="pt-3">Produk</h2>
        </div>
      </div>
      <div class="row">
        <div v-for="index in this.product_size" :key="index"
          class="col-md-6 col-xl-4 col-12 pt-3  justify-content-around d-flex">
          <ProductBox :product="products[index - 1]">
          </ProductBox>
        </div>
      </div>
    </div>
  </div>
</template>


<script>

import ProductBox from "../components/Product/ProductBox";
import CategoryBox from "../components/Category/CategoryBox";
export default {
  name: 'Home',
  components: { ProductBox, CategoryBox },
  props: ["baseURL", "products", "category"],
  data() {
    return {
      category_size: 0,
      product_size: 0
    }
  },
  mounted() {
    this.category_size = this.category.length;
    this.category_size = Math.min(4, this.category_size);

    this.product_size = this.products.length;
    this.product_size = Math.min(6, this.product_size);
  }
}
</script>

<style>
.page-holder {
  min-height: 80vh;
}

.bg-cover {
  background-size: cover !important;
}

#bg1 {
  background: url(../assets/homeview.jpg);
}

#bg2 {
  background: url(../assets/home2.jpg);
}

#bg3 {
  background: url(../assets/home3.jpg);
}

#heading {
  text-decoration: none;
  font-family: 'Roboto', sans-serif;
  font-weight: 600;
  font-size: 20px;
  opacity: 0.8;
  font-family: 'Josefin Sans', sans-serif;
  border-radius: 30px;
}

#content {
  opacity: 0.8;
}

h2 {
  font-family: 'Josefin Sans', sans-serif;
}

.fixed-button-container {
  position: absolute;
  top: 20%;
  left: 10%;
  transform: translate(-50%, -50%);
  z-index: 10;
}

.fixed-button {
  text-decoration: none;
  font-family: 'Roboto', sans-serif;
  font-weight: 400;
  opacity: 0.8;
  font-family: 'Josefin Sans', sans-serif;
  border-radius: 30px;
}
</style>
