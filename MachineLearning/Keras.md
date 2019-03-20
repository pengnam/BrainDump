# Keras

The core data structure of Keras is a **model**, which is used organise layers.


General steps to training:
1. Create model (by adding layers)
2. Compile model to configure learning process `.compile`
3. Fit model by iterating on training data in batches, or manually train
4. Evaluate performance, or generate predictions `.evaluate`, `predict`


## Model Types
Sequential Model
Linear stack of layers. 

1. Layers

Created by passing into constructor or by adding layers using the `.add` method. 

2. Specifying the Input Shape

The model needs to know the input shape that it should expect. 

3. Compilation

Receives three arguments: an optimizer, a loss- function, list of metrics


## Layers
### About
All Keras layers have a number of common methods

1. `.get_weights`: return the weights of the layer as a list of Numpy arrays
2. `.set_weights`: sets the weights of the layer from a list of Numpy arrays
3. `.get_config`: returns a dictionary containing the configuration of the layer


