<template>
  <div class="municipality-wrapper">
    <div v-for="subarea in areaCodes" :key="subarea">
      {{ subarea.name }}
      <div v-for="area in subarea.list" :key="area">
        <li>{{ areacode }}<router-link :to="`municipality/${area.area_code}`">{{ area.municipality }}</router-link></li>
      </div>
     </div>
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