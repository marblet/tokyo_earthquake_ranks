<template>
  <div class="home">
    <Title />
    <AddressForm />
  </div>
</template>

<script>
// @ is an alias to /src
import AddressForm from '@/components/AddressForm.vue'
import Title from '@/components/Title.vue'

export default {
  name: 'Home',
  components: {
    AddressForm,
    Title,
  }
}
</script>
