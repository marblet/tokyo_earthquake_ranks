<template>
  <div>
    <h2>{{ area.municipality + area.town_name }}</h2>
    <p>地盤分類：{{ area.base_class }}</p>
    <v-simple-table>
      <thead>
        <td></td>
        <td>危険量</td>
        <td>順位</td>
        <td>ランク</td>
      </thead>
      <tbody>
        <td>建物倒壊</td>
        <td>{{ area.collapse_ha }} 棟/ha</td>
        <td>{{ area.collapse_order }} / 5177</td>
        <td>{{ area.collapse_rank }}</td>
      </tbody>
      <tbody>
        <td>火災</td>
        <td>{{ area.fire_ha }} 棟/ha</td>
        <td>{{ area.fire_order }} / 5177</td>
        <td>{{ area.fire_rank }}</td>
      </tbody>
      <tbody>
        <td>活動困難</td>
        <td>{{ area.difficulty }}</td>
        <td>{{ area.difficulty_order }} / 5177</td>
        <td>{{ area.difficulty_rank }}</td>
      </tbody>
      <tbody>
        <td>総合</td>
        <td>{{ area.total_ha }} 棟/ha</td>
        <td>{{ area.total_order }} / 5177</td>
        <td>{{ area.total_rank }}</td>
      </tbody>
    </v-simple-table>
  </div>
</template>

<script>
export default {
  props: {
    id: {
      type: int
    }
  },
  created () {
    // itinialize matchedAreas using API
    axios
      .get(`/api/towninfo?id=${this.id}`)
      .then(responce => this.update(responce))
  },
}
</script>

<style>

</style>