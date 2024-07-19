<template>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <div class="container">
            <router-link class="navbar-brand" :to="{ name: 'Home' }">
                <img id="logo" src="../assets/logo.png" />
            </router-link>
            <button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarNav">
                <ul class="navbar-nav ml-auto text-white mt-2 mt-lg-0">
                    <li class="nav-item">
                        <router-link class="nav-link" :to="{ name: 'AboutUs' }">Beranda</router-link>
                    </li>
                    <li class="nav-item ml-4">
                        <router-link class="nav-link" :to="{ name: 'OurBusiness' }">Bisnis Kami</router-link>
                    </li>
                    <li class="nav-item ml-4">
                        <router-link class="nav-link" :to="{ name: 'ContactUs' }">Kontak</router-link>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>