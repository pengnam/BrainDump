# React

## Rerendering files in React
React will only update parts of the DOM that have changed.
i.e. if there is a change to the props

## React life cycles

![alt text](/Users/anton/Documents/BrainDump/WebProgramming/React-Life-Cycle.png)


`componentDidMount`: Runs after component rendered on DOM. Usually ran once

`componentWillUnmount`: Runs when the component is removed from DOM. Ran once.

`componentDidUpdate`: (After rendering) Runs when prop/state is changed

## Handling Events
React events are named using camelCase rather than lowercase. JSX passes function to event handler.

Use `e.preventDefault()` to prevent default behaviour.

React uses synthetic event. Synthetic event is a cross browser wrapper over the browser's native event.

events are not persisted in code.
