<template>
  <div class="card h-100 w-100">
    <div class="embed-responsive embed-responsive-16by9">
      <router-link v-if="!isOutOfStock" @click.native :to="{ name: 'ShowDetails', params: { id: product.id } }">
        <img class="card-img-top embed-responsive-item" :src="getImagePath(product.image)" alt="Product Image" />
      </router-link>
      <img v-else class="card-img-top embed-responsive-item out-of-stock" :src="getImagePath(product.image)"
        alt="Product Image" />
    </div>
    <div class="card-body">
      <router-link v-if="!isOutOfStock" @click.native :to="{ name: 'ShowDetails', params: { id: product.id } }">
        <h5 class="card-title">{{ product.name }}</h5>
      </router-link>
      <h5 v-else class="card-title text-muted">{{ product.name }}</h5>
      <p class="card-text"><sup>Rp</sup>{{ product.price }}</p>
      <p class="card-text font-italic">{{ strippedDescription }}...</p>
      <router-link id="edit-product" v-if="$route.name == 'AdminProduct'" @click.native
        :to="{ name: 'EditProduct', params: { id: product.id } }">
        Ubah
      </router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: "ProductBox",
  props: ["product"],
  computed: {
    strippedDescription() {
      return this.stripHtml(this.product.description).substring(0, 65);
    },
    isOutOfStock() {
      return this.product.quantity === 0 || this.product.quantity === null;
    }
  },
  methods: {
    getImagePath(image) {
      return require(`../../assets/img-main/${image}`);
    },
    stripHtml(html) {
      const div = document.createElement("div");
      div.innerHTML = html;
      return div.textContent || div.innerText || "";
    }
  }
};
</script>

<style scoped>
.embed-responsive .card-img-top {
  object-fit: cover;
}

.out-of-stock {
  filter: grayscale(100%) !important;
}

a {
  text-decoration: none;
}

.card-title {
  color: #484848;
  font-size: 1.1rem;
  font-weight: 400;
}

.card-title:hover:not(.text-muted) {
  font-weight: bold;
}

.card-text {
  font-size: 0.9rem;
}

#edit-product {
  float: right;
}
</style>