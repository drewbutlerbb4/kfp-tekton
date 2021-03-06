# Hyperparameter tuning using Katib

[Katib](https://github.com/kubeflow/katib) is a Kubernetes-based system for Hyperparameter Tuning and Neural Architecture Search. This pipeline demonstrates how to use the Katib component to find the hyperparameter tuning results. 

This pipeline uses the [MNIST dataset](http://yann.lecun.com/exdb/mnist/) to train the model and try to find the best hyperparameters using random search. Learning rate, number of convolutional layers, and the optimizer are the training parameters we want to search. Once the pipeline is completed, the ending step will print out the best hyperparameters in this pipeline experiment and clean up the workspace.

## Prerequisites 
- Install [KFP Tekton prerequisites](/samples/README.md)

## Instructions

1. First, go to the Kubeflow dashboard and create a user namespace. The Kubeflow dashboard is the endpoint to your istio-ingressgateway. We will be using the namespace `anonymous` for this example.

2. Compile the Katib pipeline. The kfp-tekton SDK will produce a Tekton pipeline yaml definition in the same directory called `katib.yaml`.
```shell
python katib.py
```

3. Next, upload the `katib.yaml` file to the Kubeflow pipeline dashboard with Tekton Backend to run this pipeline. This pipeline will run for 10 to 15 minutes.

## Acknowledgements

Thanks [Hougang Liu](https://github.com/hougangliu) for creating the original katib example.
