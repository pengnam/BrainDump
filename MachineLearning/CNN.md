#Convolutional Neural Networks

## Description
- Class of deep neural networks, typically applied to analysing imagery.
- Variation of MLP designed to require minimal preprocessing
- Able to capture spatial and temporal dependencies in an image through application of relevant features

Convolutional layers apply a convolutional operation to the input, passing the result
to the next layer.

## Layers
1. Convolutional
Each CNN processes data only for its receptive field. Describes the number of features.
--> Produces one-depth channel Convoluted Feature Output.

The output after the convolutional layer is calculated by multiplicating the weights in filter with receptive input, and then summing them.

Stride: size of steps
Padding: padding the edges with 0s

2. Pooling
Include local or global pooling layers which combine the outputs of neuron clusters
at one layer into a single neuron in the next layer.
--> Two types: max pooling, which is known to discard noisy activations, and also performs denoising, and average pooling, performs dimensionality reduction as a noise suppressing mechanism.

3. Fully Connected
Connect every neuron in one layer to every neuron in another layer.
In principle, same as MLP.

4. Receptive Field
Receives input from some number of locations in previous layer. In fully connected layer, each neuron receives input from a restricted subarea.

5. ReLU
Applying nonlinear layer introduced nonlinearity to a system. Takes only max.

## Convolution
In mathematics (in particular, functional analysis) convolution is a mathematical operation on two functions (f and g) to produce a third function that expresses how the shape of one is modified by the other.

## Architectures
- LeNet:
Conv filters applied 5x5, stride 1. Subsampling layers were 2*2 applied at stride 2.
CONV-POOL-CONV-POOL-FC-FC

- AlexNet:
4 Conv Layers
Alot of data augmentation
First time relu was applied

- VGGNet:
Smaller filters (3x3), deeper networks
    - Use smaller filters with more depth, and the stack of 3x3 conv layers have the same
effective receptive field as one 7x7 conv layer
    - fewer parameters with a smaller filter, deeper depth





GoogLeNet:

ResNet:
152 layer network

Uses the idea of a residual network.

What happens when you continue stacking deeper layers on a "plain" convolutional network.

Deeper models do not have an overfitting problem, hence it is likely an optimization problem.

Use network layers to fit a residual mapping instead of directly trying to fit a desired underlying mapping.

Use layers to fit residual F(x) + x instead of H(x) directly, where x is an output from a previous layer. The intuition is that we are training to decide the additional values in this layer.

One by one conv filter to project the depth down.
