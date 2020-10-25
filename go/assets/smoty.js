new Vue({
    el: '#null_check',
    data: {
      name: '',
      pass: ''
    },
    computed: {
      canEnter1: function() {
        if(this.name !== '') {
          return true
        } else {
          return false
        }
      }
    },
  })

  new Vue({
    el: '#app5',
    data: {
      name: '',
      pass: ''
    },
    computed: {
      canEnter1: function() {
        if(this.name !== '') {
          return true
        } else {
          return false
        }
      }
    },
  })