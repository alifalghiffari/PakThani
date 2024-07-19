<template>
  <div class="adminPanel container">
    <div class="row">
      <div class="col">
        <router-link :to="{ name: 'AdminProduct' }">
          <button type="button" class="btn btn-primary btn-lg">Produk</button>
        </router-link>
      </div>
      <div class="col">
        <router-link :to="{ name: 'AdminCategory' }">
          <button type="button" class="btn btn-primary btn-lg">Kategori</button>
        </router-link>
      </div>
      <div class="col">
        <router-link :to="{ name: 'OrderAdmin' }">
          <button type="button" class="btn btn-primary btn-lg">Pesanan Masuk</button>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  mounted() {
    if (!localStorage.getItem('token')) {
      this.$router.push({ name: 'Signin' });
    }
  }
}
</script>


<style>
.adminPanel {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.adminPanel button {
  font-size: 30px;
  width: fit-content;
  margin: 55px 0;
}
</style>