<template>
  <div class="home">
    <Title />
    <AreaList />
  </div>
</template>

<script>
// @ is an alias to /src
import AreaList from '@/components/AreaList.vue'
import Title from '@/components/Title.vue'

export default {
  name: 'Home',
  components: {
    AreaList,
    Title
  }
}
</script>
