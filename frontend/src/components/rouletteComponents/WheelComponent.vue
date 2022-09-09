<template>
  <div>
    <div class="roulette-wrapper">
      <div class="selector"></div>
      <div class="wheel" :style="isSpinning ? spinningStyle : resetStyle">
        <RowComponent v-for="n in 29" :key="n" />
      </div>
    </div>
  </div>
</template>
<script>
import RowComponent from "./RowComponent.vue";
export default {
  components: {
    RowComponent,
  },
  props: {
    isSpinning: Boolean,
    outcome: Number,
  },
  watch: {
    isSpinning: function () {
      if (this.isSpinning) {
        this.spinWheel(this.outcome);
      }
    },
  },
  data() {
    return {
      order: [0, 11, 5, 10, 6, 9, 7, 8, 1, 14, 2, 13, 3, 12, 4],
      spinningStyle: {
        "transition-timing-function": "",
        "transition-duration": "6s",
        transform: "",
      },
      resetStyle: {
        transform: "",
      },
    };
  },
  methods: {
    spinWheel(outcome) {
      var position = this.order.indexOf(outcome);

      //determine position where to land
      var rows = 12,
        card = 75 + 3 * 2,
        landingPosition = rows * 15 * card + position * card;

      var randomize = Math.floor(Math.random() * 75) - 75 / 2;

      landingPosition = landingPosition + randomize;

      var object = {
        x: Math.floor(Math.random() * 50) / 100,
        y: Math.floor(Math.random() * 20) / 100,
      };

      this.spinningStyle["transition-timing-function"] =
        "cubic-bezier(0," + object.x + "," + object.y + ",1)";
      this.spinningStyle.transform =
        "translate3d(-" + landingPosition + "px, 0px, 0px)";

      var resetTo = -(position * card + randomize);

      this.resetStyle.transform = "translate3d(" + resetTo + "px, 0px, 0px)";
    },
  },
};
</script>
<style scoped lang="scss">
.roulette-wrapper {
  position: relative;
  display: flex;
  justify-content: center;
  width: 100%;
  margin: 0 auto;
  overflow: hidden;
}

.roulette-wrapper .selector {
  width: 3px;
  background: grey;
  left: 50%;
  height: 100%;
  transform: translate(-50%, 0%);
  position: absolute;
  z-index: 2;
}

.roulette-wrapper .wheel {
  display: flex;
}
</style>
