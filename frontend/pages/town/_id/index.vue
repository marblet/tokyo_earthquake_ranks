<template>
  <div class="towninfo-wrapper">
    <div>
      <h1>{{ town.municipality + town.town_name }}</h1>
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
    <div>
      <h2>建物倒壊</h2>
      <p>{{ description.collapse.overview }}</p>
    </div>
    <div>
      <h2>火災</h2>
      <p>{{ description.fire.overview }}</p>
    </div>
    <div>
      <h2>活動困難</h2>
      <p>{{ description.difficulty.overview }}</p>
    </div>
    <div>
      <h2>総合</h2>
      <p>{{ description.total.overview }}</p>
    </div>
  </div>
</template>

<script>
import axios from 'axios'
import Meta from '@/assets/mixins/Meta.js'
import desc from '@/assets/jsons/description.json'

export default {
  name: 'TownInfo',
  mixins: [Meta],
  async asyncData ({ params, error }) {
    // should be fixed in the future
    const url = process.client ? `/api/towninfo/${params.id}` : `https://tokyo-earthquake-ranks.df.r.appspot.com/api/towninfo/${params.id}`
    return axios
      .get(url)
      .then(response => {
        if (response.data.Error) {
          return error({
            statusCode: 404
          })
        }
        return {
          town: response.data
        }
    })
  },
  head () {
    return {
      title: this.town.municipality + this.town.town_name
    }
  },
  data () {
    return {
      town: {},
      meta: {
        // title: town.municipality + town.town_name,
        description: 'Towninfo page',
        type: 'article',
        url: 'https://example.com/test',
        // image: 'https://example.com/img/ogp/test.jpg',
      },
      description: desc
    }
  },
  methods: {
    update (response) {
      this.town = response.data
    }
  },
}
</script>

<style>
@import '@/assets/styles/loader.css';
.towninfo-wrapper {
  margin: auto;
  width: 100%;
  max-width: 1200px;
}
</style>