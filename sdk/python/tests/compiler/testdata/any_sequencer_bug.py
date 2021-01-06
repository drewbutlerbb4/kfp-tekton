# Copyright 2020 kubeflow.org
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from kfp import dsl
from kfp_tekton.compiler import TektonCompiler
from kfp_tekton.compiler.any_sequencer import after_any

class FlipCoinOp(dsl.ContainerOp):

    def __init__(self, name, forced_result=""):
        super(FlipCoinOp, self).__init__(
            name=name,
            image='python:alpine3.6',
            command=['sh', '-c'],
            arguments=['exit(1)'],
            file_outputs={'output_value': '/tmp/output'})

class PrintOp(dsl.ContainerOp):

    def __init__(self, name, msg):
        super(PrintOp, self).__init__(
            name=name,
            image='alpine:3.6',
            command=['echo', msg])

@dsl.pipeline(
    name="Any Sequencer",
    description="Any Sequencer Component Demo",
)
def any_sequence_pipeline(
):
    flip = FlipCoinOp('flip', "heads")
    flip2 = FlipCoinOp('flip2', "tails")
    flip3 = FlipCoinOp('flip3', "heads")

    PrintOp('print1', flip.output).apply(after_any([flip, flip2, flip3]))

if __name__ == "__main__":
    TektonCompiler().compile(any_sequence_pipeline, "any_sequencer" + ".yaml")
