<template>
  <div>
    <client-only>
      <codemirror v-model="code" :options="options" />
    </client-only>
  </div>
</template>
<script>
  export default {
    name: 'BXml',
    data() {
      return {
        code: '',
        options: {
          lineNumbers: true,
          matchBrackets: true,
          mode: 'application/xml',
          indentUnit: 4,
          indentWithTabs: true,
          autoCloseTags: true,
          moveOnDrag: false
        }
      }
    }
  }

</script>
