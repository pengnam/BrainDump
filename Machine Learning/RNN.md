# Residual Neural Network


# Description
ANN builds pyrimidal cells in the cerebral cortex. Uses shortcuts to jump over layers.

Motivation for skipping over layers is to avoid problem of vanishing gradients by reusing activations from a previous layer until adjacent layer has learnt weights.

The weights adapt to mute the upstream layer, amplify the previously skipped layer.

Forward propagation for a simple RNN:
![alttext](https://wikimedia.org/api/rest_v1/media/math/render/svg/380c8bb2baa89520876ac59d05b740562ad7814f)
