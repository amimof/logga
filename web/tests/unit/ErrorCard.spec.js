import { shallowMount } from '@vue/test-utils'
import ErrorCard from '@/components/ErrorCard.vue'

describe('ErrorCard.vue', () => {
  it('renders props.title and props.error when passed', () => {
    const title = 'Something happened'
    let error = new Error("Unknown error")
    const wrapper = shallowMount(ErrorCard, {
      propsData: { title, error }
    })
    expect(wrapper.text()).toMatch(title)
  })
})
