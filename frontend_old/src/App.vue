<template>
  <v-app>
    <Header />
    <v-content>
      <div class="main-content">
        <router-view/>
      </div>
    </v-content>
    <Footer />
  </v-app>
</template>

<script>
import Footer from './components/Footer';
import Header from './components/Header';

export default {
  name: 'App',

  components: {
    Footer,
    Header,
  },

  data: () => ({
    //
  }),
};
</script>

<style lang="scss">
.main-content {
  width: 1100px;
  margin: auto;
  padding: {
    bottom: 10px;
    left: 5px;
    right: 5px;
    top: 10px;
  }
}
</style>
