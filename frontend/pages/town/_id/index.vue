<template>
  <div>
    <div v-show="loading" class="loader">
      <p>loading</p>
    </div>
    <div v-show="!loading">
      <h2>{{ town.municipality + town.town_name }}</h2>
      <p>地盤分類：{{ town.base_class }}</p>
      <v-simple-table>
        <thead>
          <tr>
            <th></th>
            <th>危険量</th>
            <th>順位</th>
            <th>ランク</th>
          </tr>
        </thead>
        <tbody>
          <tr>
            <td>建物倒壊</td>
            <td>{{ town.collapse_ha }} 棟/ha</td>
            <td>{{ town.collapse_order }} / 5177</td>
            <td>{{ town.collapse_rank }}</td>
          </tr>
          <tr>
            <td>火災</td>
            <td>{{ town.fire_ha }} 棟/ha</td>
            <td>{{ town.fire_order }} / 5177</td>
            <td>{{ town.fire_rank }}</td>
          </tr>
          <tr>
            <td>活動困難</td>
            <td>{{ town.difficulty }}</td>
            <td>{{ town.difficulty_order }} / 5177</td>
            <td>{{ town.difficulty_rank }}</td>
          </tr>
          <tr>
            <td>総合</td>
            <td>{{ town.total_ha }} 棟/ha</td>
            <td>{{ town.total_order }} / 5177</td>
            <td>{{ town.total_rank }}</td>
          </tr>
        </tbody>
      </v-simple-table>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
  data () {
    return {
      loading: true,
      town: {'municipality': '', 'town_name': ''},
    }
  },
  created () {
    // itinialize matchedTowns using API
    axios
      .get(`/api/towninfo?id=${this.$route.params['id']}`)
      .then(responce => this.update(responce))
  },
  methods: {
    update (responce) {
      this.town = responce.data
      this.loading = false
    }
  },
  head () {
    return {
      title: this.town.municipality + this.town.town_name
    }
  },
}
</script>

<style>
@import '@/assets/styles/loader.css'
</style>