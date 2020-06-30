<template>
  <div>
    <TownList />
  </div>
</template>

<script>
import TownList from '@/components/TownList.vue'

export default {
  name: 'MunicipalityList',
  components: {
    TownList
  },
  async asyncData ({ params, error }) {
    // should be fixed in the future
    const url = process.client ? `/api/towninfo?id=${params.area}` : `https://tokyo-earthquake-ranks.df.r.appspot.com/api/towninfo?id=${params.id}`
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
}
</script>