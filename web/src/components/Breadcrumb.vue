<template>
  <b-breadcrumb :items="items" />
</template>

<script>
export default {
  name: 'Breadcrumb',
  computed: {
    items() {
      let bc = this.$route.matched;
      let res = [];
      switch(bc[0].name) {
        case 'Namespaces':
          res.push({
            text: 'namespaces',
            to: { name: 'Namespaces' }
          });
          break;
        case 'Pods':
          res.push({
            text: 'namespaces',
            to: { name: 'Namespaces' }
          })
          res.push({
            text: this.$route.params.namespace,
            to: { name: 'Pods', namespace: this.$route.params.namespace }
          })
          break;
        case 'Pod':
          res.push({
            text: 'namespaces',
            to: { name: 'Namespaces' }
          })
          res.push({
            text: this.$route.params.namespace,
            to: { name: 'Pods', namespace: this.$route.params.namespace }
          })
          res.push({
            text: this.$route.params.pod,
            to: { name: 'Pod', pod: this.$route.params.pod }
          })
          break;
      }
      return res
    }
  }
}
</script>

<style lang="scss" scoped>
</style>


