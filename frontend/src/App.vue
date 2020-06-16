<template>
  <v-app>
    <v-content>
      <AreaList/>
    </v-content>
  </v-app>
</template>

<script>
import AreaList from './components/AreaList';

export default {
  name: 'App',

  components: {
    AreaList,
  },

  data: () => ({
    //
  }),
};
</script>
