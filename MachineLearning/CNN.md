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

    - Parameter sharing
Parameter sharing scheme is used in Convolutional Layers to control the number of parameters. Using the real-world example above, we see that there are 55*55*96 = 290,400 neurons in the first Conv Layer, and each has 11*11*3 = 363 weights and 1 bias. Together, this adds up to 290400 * 364 = 105,705,600 parameters on the first layer of the ConvNet alone. Clearly, this number is very high.

It turns out that we can dramatically reduce the number of parameters by making one reasonable assumption: That if one feature is useful to compute at some spatial position (x,y), then it should also be useful to compute at a different position (x2,y2). In other words, denoting a single 2-dimensional slice of depth as a depth slice (e.g. a volume of size [55x55x96] has 96 depth slices, each of size [55x55]), we are going to constrain the neurons in each depth slice to use the same weights and bias. With this parameter sharing scheme, the first Conv Layer in our example would now have only 96 unique set of weights (one for each depth slice), for a total of 96*11*11*3 = 34,848 unique weights, or 34,944 parameters (+96 biases). Alternatively, all 55*55 neurons in each depth slice will now be using the same parameters. In practice during backpropagation, every neuron in the volume will compute the gradient for its weights, but these gradients will be added up across each depth slice and only update a single set of weights per slice.

 Notice that the parameter sharing assumption is relatively reasonable: If detecting a horizontal edge is important at some location in the image, it should intuitively be useful at some other location as well due to the translationally-invariant structure of images. There is therefore no need to relearn to detect a horizontal edge at every one of the 55*55 distinct locations in the Conv layer output volume.


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


## Visualisation
Question:
What are the intermediate features looking for (intermediate layers)?


First layer:
Visualize Filters
-> Humans tend to visualize/represent oriented edges, opposing colours.
(regardless or architecture)

Visualizing weights tend to visualize what the weight is looking for because
the product of two vectors produces a maximal output.


Last Layers:
Visualise the nearest neighbors to the test set features
--> The nearest value will be in terms of feature space instead of exact pixels

Use a t-SNE dimensional feature space to compress the dimensions to 2
Show the clustering of images


Intermediate:
Visualize the activations
--> Allows you to make guesses about the features detected by each layer.


Occlusion:
Mask part of the image before feeding to CNN, draw heatmap of probability at each mask location.

Retrieving images that maximally activate a neuron





Links:
https://github.com/keras-team/keras/blob/master/examples/conv_filter_visualization.py
https://lvdmaaten.github.io/tsne/
