<template>
  <div class="municipality-wrapper">
    <li v-for="areacode in areaCodes" :key="">{{ areacode }}</li>
    <p>23区</p>
    <p>23区外</p>
  </div>
</template>

<script>
import Areacodes from '@/assets/jsons/areacodes.json'

export default {
  name: 'Municipality',
  data () {
    return {
      areaCodes: Areacodes
    }
  }
}
</script>