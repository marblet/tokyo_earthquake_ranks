<template>
  <div class="home">
    <Title />
    <TownList />
  </div>
</template>

<script>
// @ is an alias to /src
import Title from '@/components/Title.vue'
import TownList from '@/components/TownList.vue'

export default {
  name: 'Home',
  components: {
    Title,
    TownList,
  }
}
</script>
