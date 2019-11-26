import { Card, Col, message, Row, Statistic } from "antd";
import _ from "lodash";
import { ModelzooServicePromiseClient } from "protos/services_grpc_web_pb";
import { Empty } from "protos/services_pb";
import React, { FC, useMemo, useState } from "react";
import { FeaturedModelTag } from "../Components/Tags";
import { ModelObject, parseModels, TestObject, parseTests } from "../Utils/ProtoUtil";

import { Link } from "react-router-dom";
import { TagsSet, StatsSet } from "../Components/Tags";

interface HomeProps {
  models: ModelObject[];
  client: ModelzooServicePromiseClient;
  filter: String;
}

interface CatalogProps {
  // models: ModelObject[];
  client: ModelzooServicePromiseClient;
  filter: String;
}

export const Catalog: FC<CatalogProps> = props => {
  // let { models } = props;
  let { client, filter } = props;
  const [models, setModels] = useState<Array<ModelObject>>([]);
  const [tests, setTests] = useState<Array<TestObject>>([]);

  useMemo(() => {
    client
      .listModels(new Empty(), undefined)
      .then(resp => setModels(parseModels(resp.getModelsList())))
      .catch(err => message.error("Unable to fetch all models"));
  }, [client]);

  useMemo(() => {
    client
      .listTests(new Empty(), undefined)
      .then(resp => setTests(parseTests(resp.getTestsList())))
      .catch(err => message.error("Unable to fetch all tests"));
  }, [client]);

  let cards = models.map((model: ModelObject, index, arr) => {
    var show = false;
    var filter_values = filter.split(" ");
    filter_values = filter_values.filter(function(e) {
      return e.trim().length != 0;
    })
    show = show || filter.length == 0;
    for (var i = 0; i < filter_values.length; i++) {
      show = show || model.name.toLowerCase().includes(filter_values[i].toLowerCase());
    }
    Object.keys(model.metadata).map(function(key) {
      for (var i = 0; i < filter_values.length; i++) {
        show = show || model.metadata[key][0].toLowerCase().includes(filter_values[i].toLowerCase());
      }
    })
    console.log(model.metadata)
    if (show) {
      return (
        <Col span={8}>
          <Card
            title={model.name}
            style={{ margin: "2px" }}
            extra={<Link to={`model/${model.name}`}>Test it</Link>}
          >
            <Row>
              <StatsSet model={model} showAll={false}></StatsSet>
            </Row>
            <TagsSet model={model} showAll={false}></TagsSet>
          </Card>
        </Col>
      );
    }
  });

  let testcards = tests.map((test: TestObject, index, arr) => {
    var show = false;
    var filter_values = filter.split(" ");
    filter_values = filter_values.filter(function(e) {
      return e.trim().length != 0;
    })
    show = show || filter.length == 0;
    for (var i = 0; i < filter_values.length; i++) {
      show = show || test.name.toLowerCase().includes(filter_values[i].toLowerCase());
    }
    Object.keys(test.metadata).map(function(key) {
      for (var i = 0; i < filter_values.length; i++) {
        show = show || test.metadata[key][0].toLowerCase().includes(filter_values[i].toLowerCase());
      }
    })
    console.log(test.metadata)
    if (show) {
      return (
        <Col span={8}>
          <Card
            title={test.name}
            style={{ margin: "2px" }}
          >
            <Row>
              <StatsSet model={test} showAll={false}></StatsSet>
            </Row>
            <TagsSet model={test} showAll={false}></TagsSet>
          </Card>
        </Col>
      );
    }
  });

  return (
    <div>
      <Row>{cards}</Row>
      <Row>{testcards}</Row>
    </div>
  );
};

export const Home: FC<HomeProps> = props => {
  // let { models } = props;
  let { client, filter } = props;
  return <Catalog client={client} filter={filter}></Catalog>;
};
