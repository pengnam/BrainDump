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

2. Pooling
Include local or global pooling layers which combine the outputs of neuron clusters
at one layer into a single neuron in the next layer.
--> Two types: max pooling, which is known to discard noisy activations, and also performs denoising, and average pooling, performs dimensionality reduction as a noise suppressing mechanism.

3. Fully Connected
Connect every neuron in one layer to every neuron in another layer.
In principle, same as MLP.

4. Receptive Field
Receives input from some number of locations in previous layer. In fully connected layer, each neuron receives input from a restricted subarea.

## Convolution
In mathematics (in particular, functional analysis) convolution is a mathematical operation on two functions (f and g) to produce a third function that expresses how the shape of one is modified by the other.
## Architectures
LeNet
AlexNet
VGGNet
GoogLeNet
ResNet
ZFNet
