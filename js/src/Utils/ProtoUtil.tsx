import { Model as pb_Model, Test as pb_Test, KVPair } from "protos/services_pb";
import _ from "lodash";

export interface ModelObject {
  name: string;
  metadata: Record<string, string[]>;
}

export interface TestObject {
  name: string;
  metadata: Record<string, string[]>;
}

function metadataToKV(pairs: KVPair[]): Record<string, string[]> {
  return _.mapValues(
    _.groupBy(pairs, (pair: KVPair) => pair.getKey()),
    (lst: KVPair[]) => {
      return _.map(lst, pair => pair.getValue());
    }
  );
}

export const parseModels = (models: Array<pb_Model>): Array<ModelObject> => {
  return models.map((model: pb_Model, index, array) => {
    return {
      name: model.getModelName(),
      metadata: metadataToKV(model.getMetadataList())
    };
  });
};

export const parseTests = (tests: Array<pb_Test>): Array<TestObject> => {
  return tests.map((test: pb_Test, index, array) => {
    return {
      name: test.getTestName(),
      metadata: metadataToKV(test.getMetadataList())
    };
  });
};
