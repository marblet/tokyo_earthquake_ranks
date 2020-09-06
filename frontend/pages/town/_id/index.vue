<template>
  <div class="towninfo-wrapper">
    <div>
      <div class="towntitle">
        <h1>{{ town.municipality + town.town_name }}</h1>
        <p><nuxt-link :to="`/municipality/${town.area_code}`">{{ town.municipality }}</nuxt-link> > {{ town.town_name }}</p>
      </div>
      <p>地盤分類：{{ town.base_class }}</p>
      <v-card>
        <v-card-title>総合ランク: {{ town.total_rank }}</v-card-title>
        <v-card-text>
          <p>危険量: {{ town.total_ha }} 棟/ha</p>
          <p>順位: {{ town.total_order }} / 5177</p>
        </v-card-text>
      </v-card>
      <v-container fluid>
        <v-row dense>
          <v-col>
            <v-card>
              <v-card-title>建物倒壊: {{ town.collapse_rank }}</v-card-title>
              <v-card-text>
                <p>危険量: {{ town.collapse_ha }} 棟/ha</p>
                <p>順位: {{ town.collapse_order }} / 5177</p>
              </v-card-text>
            </v-card>
          </v-col>
          <v-col>
            <v-card>
              <v-card-title>火災: {{ town.fire_rank }}</v-card-title>
              <v-card-text>
                <p>危険量: {{ town.fire_ha }} 棟/ha</p>
                <p>順位: {{ town.fire_order }} / 5177</p>
              </v-card-text>
            </v-card>
          </v-col>
          <v-col>
            <v-card>
              <v-card-title>活動困難: {{ town.difficulty_rank }}</v-card-title>
              <v-card-text>
                <p>困難度: {{ town.difficulty }}</p>
                <p>順位: {{ town.difficulty_order }} / 5177</p>
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
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
    const url = process.client ? `/api/towns/${params.id}` : `https://tokyo-earthquake-ranks.df.r.appspot.com/api/towns/${params.id}`
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
.towntitle {
  border-bottom: 4px solid black;
  display: flex;
  justify-content: space-between;
  align-items: flex-end;
}
</style>