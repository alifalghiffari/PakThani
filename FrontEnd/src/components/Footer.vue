<template>
  <footer>
    <div class="container pt-5">
      <div class="row">
        <div class="col-12">
          <div class="horizontal-line"></div>
        </div>
      </div>
      <div class="row align-items-center justify-content-between pt-5 pb-5">
        <div class="col-md-4 text-center">
          <img id="logo" src="../assets/logo.png" alt="Pakthani Logo" />
        </div>
        <div class="col-md-4 text-start">
          <div class="contact-info">
            <p class="mb-3 d-flex justify-content-start">
              <ion-icon name="location-outline"></ion-icon>
              Jl. Kuta Lampanah, Kuta Alam, Banda Aceh
            </p>
            <p class="mb-3 d-flex justify-content-start">
              <ion-icon name="call-outline"></ion-icon>
              +62 822-7220-7548
            </p>
            <p class="mb-2 d-flex justify-content-start">
              <a href="https://www.instagram.com/pakthani.id" target="_blank"
                class="d-flex align-items-center text-decoration-none" style="color: inherit;">
                <ion-icon name="logo-instagram" class="me-2"></ion-icon>
                pakthani.id
              </a>
            </p>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-12">
          <div class="horizontal-line"></div>
        </div>
      </div>
      <div class="row pt-3">
        <div class="col-8 d-flex justify-content-start">
          <router-link class="footer-link" @click.native="" :to="{ name: 'AboutUs' }">Tentang Kami</router-link>
          <a class="footer-link ms-3"
            href="https://api.whatsapp.com/send?phone=6282272207548&text=Selamat%20datang%20di%20Pakthani,%20your%20online%20grocery%20anywhere"
            target="_blank">Kontak Kami</a>
        </div>
        <div class="col d-flex justify-content-end">
          <p>Copyright © 2023 • PakThani</p>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: "Footer"
}
</script>

<style>
footer {
  margin-top: 100px;
  background-color: #232F3E;
  font-size: 16px;
  color: #DDD;
  padding: 40px 0;
}

#logo {
  width: 150px;
}

.contact-info {
  color: #DDD;
}

.contact-info p {
  margin: 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.contact-info ion-icon {
  margin-right: 10px;
  font-size: 24px;
}

.horizontal-line {
  border-top: 1px solid #DDD;
  margin: 0 0 20px;
}

.footer-link {
  color: #DDD;
  font-size: 18px;
  text-decoration: none;
  margin-right: 20px;
}

.footer-link:hover {
  text-decoration: underline;
}

.text-center p {
  margin: 0;
}

.end {
  justify-content: end !important;
}
</style>
