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
import Meta from '@/assets/mixins/Meta.js'

export default {
  name: 'TownInfo',
  mixins: [Meta],
  data () {
    return {
      loading: true,
      town: {'municipality': '', 'town_name': ''},
      meta: {
        title: this.$route.params['id'],
        description: 'Towninfo page',
        type: 'article',
        url: 'https://example.com/test',
        // image: 'https://example.com/img/ogp/test.jpg',
      }
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
}
</script>

<style>
@import '@/assets/styles/loader.css'
</style>