<template>
  <div class="card h-100 w-100">
    <div class="embed-responsive embed-responsive-16by9">
      <router-link @click.native :to="{ name: 'ListProducts', params: { id: category.id } }">
        <img
          class="card-img-top embed-responsive-item"
          :src="getImagePath(category.icon)"
          alt="Category Image"
        />
      </router-link>
    </div>

    <div class="card-body">
      <router-link @click.native :to="{ name: 'ListProducts', params: { id: category.id } }"
        ><h5 class="card-title">{{ category.category }}</h5></router-link
      >
      <router-link
        id="edit-category"
        @click.native :to="{ name: 'EditCategory', params: { id: category.id } }"
        v-show="$route.name == 'AdminCategory'"
      >
        Edit
      </router-link>
    </div>
  </div>
</template>

<script>
export default {
  name: "CategoryBox",
  props: ["category"],
  methods: {
    getImagePath(image) {
      return require(`../../assets/Icon/${image}`);
    },
    listProducts() {
      this.$router.push({
        name: "ListProducts",
        params: { id: this.category.id },
      });
    },
  },
};
</script>

<style scoped>
.embed-responsive {
  display: flex;
  align-items: center;
  justify-content: center;
}

.card-img-top {
  padding-left: 10%; 
  width: 80%; 
  height: auto;
  object-fit: cover;
}

a {
  text-decoration: none;
}

.card-title {
  color: #484848;
  font-size: 1.1rem;
  font-weight: 400;
}

.card-title:hover {
  font-weight: bold;
}

.card-text {
  font-size: 0.9rem;
}

#edit-category {
  float: right;
}
</style>
