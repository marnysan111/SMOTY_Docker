new Vue({
    el: '#null_check',
    data: {
      name: "",
      password: "",
    },
    computed: {
      canEnter1: function() {
        if(this.name !== "" && this.password !== "") {
          return true
        } else {
          return false
        }
      }
    },
  })
