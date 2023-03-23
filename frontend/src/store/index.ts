import { createStore } from 'vuex'
import axios from 'axios'

export default createStore({
  state: {
    counter: 0,
    color: 'red'
  },
  getters: {
    getCounterSquare (state) {
      return state.counter * state.counter
    }
  },
  // commit
  mutations: {
    decreaseCounter (state, value) {
      state.counter -= value
    },
    increaseCounter (state, value) {
      state.counter += value
    },
    setColor (state, color) {
      state.color = color
    }
  },
  // Dispatch
  actions: {
    decreaseCounter (context, value) {
      context.commit('decreaseCounter', value)
    },
    increaseCounter (context, value) {
      context.commit('increaseCounter', value)
    },
    setColor (context, color) {
      const randomURL = 'https://www.random.org/integers/?num=1&min=1&max=6&col=1&base=10&format=plain&rnd=new'
      axios.get(randomURL).then((response) => {
        console.log(response.data)
      })
      context.commit('setColor', color)
    }
  },
  modules: {
  }
})
