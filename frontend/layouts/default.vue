<template>
  <v-app>
    <Header />
    <v-main>
      <nuxt />
    </v-main>
    <Footer />
  </v-app>
</template>

<script>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'

export default {
  components: {
    Header,
    Footer,
  }
}
</script>

<style lang="scss">
.v-main {
  margin: auto;
  width: 100%;
  max-width: 1200px;
}
</style>
