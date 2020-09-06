<template>
  <div class="top-wrapper">
    <div class="top-content-wrapper">
      <Title />
      <div class="addressform-wrapper">
        <AddressForm />
      </div>
      <Municipality />
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import AddressForm from '@/components/AddressForm.vue'
import Municipality from '@/components/Municipality.vue'
import Title from '@/components/Title.vue'

import Meta from '@/assets/mixins/Meta.js'

export default {
  name: 'Home',
  components: {
    AddressForm,
    Municipality,
    Title,
  },
  mixins: [Meta],
  data () {
    return {
      meta: {
        title: 'Find Safe City',
        description: '東京都が実施した地域調査をもとに、地震に強い街を見つけましょう。',
        type: 'article',
        url: 'https://example.com/test',
        // image: 'https://example.com/img/ogp/test.jpg',
      }
    }
  }
}
</script>

<style lang="scss">
.top-wrapper {
  background: {
    // image: url("~@/assets/3432480_m.jpg");
    image: url("https://drive.google.com/uc?export=view&id=1rp_rWp1EC5eVPAMhK9xlV3XW0e2rGFel");
    attachment: fixed;
  }
}
.top-content-wrapper {
  margin: auto;
  width: 100%;
  max-width: 1200px;
  padding: {
    top: 10px;
    bottom: 20px;
    left: 10px;
    right: 10px;
  }
}
.addressform-wrapper {
  margin: auto;
  width: 100%;
  max-width: 1000px;
  height: 200px;
}
</style>