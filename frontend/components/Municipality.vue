<template>
  <div class="municipality-wrapper">
    <p>市区町村から探す</p>
    <div v-for="subarea in areaCodes" :key="subarea.name" class="subarea-wrapper">
      <h2>{{ subarea.name }}</h2>
      <ul>
        <li v-for="area in subarea.list" :key="area.area_code"><router-link :to="`municipality/${area.area_code}`">{{ area.municipality }}</router-link></li>
      </ul>
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

<style lang="scss">
.municipality-wrapper {
  padding: {
    top: 15px;
    bottom: 15px;
    left: 15px;
    right: 15px;
  }
  background: {
    color: #FFF;
  }
}
.subarea-wrapper ul {
  overflow: hidden;
}
.subarea-wrapper li {
  float: left;
  width: 220px;
  margin: {
    bottom: 10px;
  }
}
</style>