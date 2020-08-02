const vm = new Vue ({
    el: "#app",
    data: function() {
      return {
        show: false
      }
    },
    methods: {
      display_switch: function() {
        this.show = !this.show
      }
    }
  })