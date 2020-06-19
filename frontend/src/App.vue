<template>
  <v-app>
    <Header />
    <v-content>
      <AreaList/>
    </v-content>
    <Footer />
  </v-app>
</template>

<script>
import AreaList from './components/AreaList';
import Footer from './components/Footer';
import Header from './components/Header';

export default {
  name: 'App',

  components: {
    AreaList,
    Footer,
    Header,
  },

  data: () => ({
    //
  }),
};
</script>
