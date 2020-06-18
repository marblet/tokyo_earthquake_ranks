<template>
  <v-app>
    <Header />
    <v-content>
      <AreaList/>
    </v-content>
    <v-footer>FindSafeCity</v-footer>
  </v-app>
</template>

<script>
import AreaList from './components/AreaList';
import Header from './components/Header';

export default {
  name: 'App',

  components: {
    AreaList,
    Header,
  },

  data: () => ({
    //
  }),
};
</script>
