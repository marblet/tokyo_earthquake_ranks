<template>
  <v-app>
    <Header />
    <v-content>
      <router-view/>
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
