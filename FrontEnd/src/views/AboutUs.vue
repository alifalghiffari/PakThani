<template>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <div class="container">
            <router-link class="navbar-brand" :to="{ name: 'Home' }">
                <img id="logo" src="../assets/logo.png" />
            </router-link>
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav"
                aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav ml-auto text-white mt-2 mt-lg-0">
                    <li class="nav-item">
                        <router-link class="nav-link" :to="{ name: 'AboutUs' }">Beranda</router-link>
                    </li>
                    <!-- <li class="nav-item ml-4">
                        <router-link class="nav-link" :to="{ name: 'OurBusiness' }">Bisnis Kami</router-link>
                    </li> -->
                </ul>
            </div>
        </div>
    </nav>

    <div class="page-img bg-cover" id="bg10">
        <div class="container">
            <div class="row align-items-center justify-content-center" style="min-height: 70vh;">
                <div class="col-md-9 text-center">
                    <p class="breadcrumbs text-dark">
                        <span class="mr-2">
                            <router-link class="text-dark" :to="{ name: 'Home' }">Home /</router-link>
                        </span>
                        <span>About Us</span>
                    </p>
                    <h1 class="mb-0 bread text-light">
                        <img src="../assets/logo.png" alt="About Us" class="img-fluid">
                    </h1>
                </div>
            </div>
        </div>
    </div>


    <section class="bg-light">
        <div class="container">
            <div class="row">
                <div class="col-md-5 p-md-5 img img-2 d-flex justify-content-center align-items-center" id="bg11"></div>
                <div class="col-md-7 py-5 wrap-about pb-md-5 ftco-animate fadeInUp ftco-animated">
                    <div class="heading-section-bold mb-4 mt-md-5">
                        <div class="ml-md-0">
                            <h1 class="mb-4">Selamat datang di PakThani</h1>
                        </div>
                    </div>
                    <div class="pb-md-5">
                        <p>Selamat datang di PakThani, tujuan utama Anda untuk bahan baku makanan berkualitas tinggi.
                            Kami adalah platform ecommerce yang berdedikasi untuk menyediakan bahan baku terbaik bagi
                            kebutuhan kuliner Anda, baik untuk rumah tangga maupun bisnis.</p>
                        <p>Di PakThani, misi kami adalah menghubungkan petani lokal dengan konsumen secara langsung,
                            memastikan setiap produk yang Anda terima adalah hasil dari sumber yang terpercaya dan
                            berkualitas. Kami percaya bahwa makanan yang baik dimulai dari bahan baku yang baik, dan
                            itulah yang kami hadirkan untuk Anda.</p>
                        <p>
                            <span class="mr-2">
                                <router-link class="btn btn-primary" :to="{ name: 'Home' }">Mulai
                                    Berbelanja</router-link>
                            </span>
                        </p>
                    </div>
                </div>
            </div>
        </div>
    </section>

</template>

<script>
export default {
    name: 'AboutUs',
};
</script>

<style>
.page-img {
    min-height: 70vh;
}

.bg-cover {
    background-size: cover !important;
}

#bg10 {
    background: url('../assets/vegetables.jpg') no-repeat center center;
    background-size: cover;
    height: 70vh;
}

#bg11 {
    background: url('../assets/food.jpg') no-repeat center center;
    background-size: cover;
    height: 100vh;
    margin-top: 50px;
    /* Adjust this value to control the vertical position */
}
</style>
