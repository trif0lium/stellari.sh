import * as gcp from "@pulumi/gcp";
import * as pulumi from "@pulumi/pulumi";
import { Config } from "@pulumi/pulumi";
import * as random from "@pulumi/random";

const config = new Config();

export const region = config.get("region") || "asia-southeast1";
export const nodeCount = config.getNumber("nodeCount") || 2;
export const nodeMachineType = config.get("nodeMachineType") || "n1-standard-2";
export const username = config.get("username") || "admin";
export const password = new random.RandomPassword("password", {
  length: 20,
  special: true,
}).result;

export const engineVersion = gcp.container.getEngineVersions({
  location: region,
});

export const masterVersion = engineVersion.then((it) => it.latestMasterVersion);
export const named = (name: string) => `${pulumi.getProject()}-${pulumi.getStack()}-${name}`;
