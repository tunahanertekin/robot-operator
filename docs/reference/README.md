# API Reference


## robot.roboscale.io/v1alpha1

Package v1alpha1 contains API Schema definitions for the robot v1alpha1 API group

### Resource Types
- [Robot](#robot)
- [WorkspaceManager](#workspacemanager)
- [BuildManager](#buildmanager)
- [LaunchManager](#launchmanager)
- [RobotDevSuite](#robotdevsuite)
- [MetricsExporter](#metricsexporter)
- [RobotVDI](#robotvdi)
- [RobotIDE](#robotide)
- [DiscoveryServer](#discoveryserver)
- [ROSBridge](#rosbridge)
- [RobotArtifact](#robotartifact)



#### Robot



Robot is the custom resource that contains ROS 2 components (Workloads, Cloud VDI, Cloud IDE, ROS Bridge, Configurational Resources), robolaunch Robot instances can be decomposed and distributed to both cloud instances and physical instances using federation.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `robot.roboscale.io/v1alpha1`
| `kind` _string_ | `Robot`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[RobotSpec](#robotspec)_ | Specification of the desired behavior of the Robot. |
| `status` _[RobotStatus](#robotstatus)_ | Most recently observed status of the Robot. |


#### RobotSpec



RobotSpec defines the desired state of Robot.

_Appears in:_
- [Robot](#robot)
- [RobotArtifact](#robotartifact)

| Field | Description |
| --- | --- |
| `distributions` _[ROSDistro](#rosdistro) array_ | ROS 2 distributions to be used. You can select multiple distributions if they are supported in the same underlying OS. (eg. `foxy` and `galactic` are supported in Ubuntu Focal, so they can be used together but both cannot be used with `humble`) |
| `rmwImplementation` _[RMWImplementation](#rmwimplementation)_ | RMW implementation selection. Robot operator currently supports only FastRTPS. See https://docs.ros.org/en/foxy/How-To-Guides/Working-with-multiple-RMW-implementations.html. |
| `storage` _[Storage](#storage)_ | Total storage amount to persist via Robot. Unit of measurement is MB. (eg. `10240` corresponds 10 GB) This amount is being shared between different components. |
| `discoveryServerTemplate` _[DiscoveryServerSpec](#discoveryserverspec)_ | Discovery server configurational parameters. |
| `rosBridgeTemplate` _[ROSBridgeSpec](#rosbridgespec)_ | ROS bridge configurational parameters. |
| `robotDevSuiteTemplate` _[RobotDevSuiteSpec](#robotdevsuitespec)_ | Robot development suite template |
| `workspaceManagerTemplate` _[WorkspaceManagerSpec](#workspacemanagerspec)_ | Workspace manager template to configure ROS 2 workspaces. |
| `buildManagerTemplate` _[BuildManagerSpec](#buildmanagerspec)_ | [*alpha*] Build manager template for initial configuration. |
| `launchManagerTemplates` _[LaunchManagerSpec](#launchmanagerspec) array_ | [*alpha*] Launch manager template for initial configuration. |
| `development` _boolean_ | [*alpha*] Switch to development mode if `true`. |
| `rootDNSConfig` _[RootDNSConfig](#rootdnsconfig)_ | [*alpha*] Root DNS configuration. |
| `tlsSecretRef` _[TLSSecretReference](#tlssecretreference)_ | [*alpha*] TLS secret reference. |


#### RobotStatus



RobotStatus defines the observed state of Robot.

_Appears in:_
- [Robot](#robot)

| Field | Description |
| --- | --- |
| `phase` _RobotPhase_ | Phase of Robot. It sums the general status of Robot. |
| `image` _string_ | Main image of Robot. It is derived either from the specifications or determined directly over labels. |
| `nodeName` _string_ | Node that Robot uses. It is selected via tenancy labels. |
| `volumeStatuses` _[VolumeStatuses](#volumestatuses)_ | Robot persists some of the directories of underlying OS inside persistent volumes. This field exposes persistent volume claims that dynamically provision PVs. |
| `discoveryServerStatus` _[DiscoveryServerInstanceStatus](#discoveryserverinstancestatus)_ | Discovery server instance status. |
| `rosBridgeStatus` _[ROSBridgeInstanceStatus](#rosbridgeinstancestatus)_ | ROS bridge instance status. |
| `loaderJobStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of loader job that configures environment. |
| `workspaceManagerStatus` _[WorkspaceManagerInstanceStatus](#workspacemanagerinstancestatus)_ | Workspace manager instance status if exists. |
| `robotDevSuiteStatus` _[RobotDevSuiteInstanceStatus](#robotdevsuiteinstancestatus)_ | Robot development suite instance status. |
| `attachedBuildObject` _[AttachedBuildObject](#attachedbuildobject)_ | Attached build object information. A BuildManager can be attached with a label on it with key `robolaunch.io/target-robot` and value of the target robot's name. Robot sorts the BuildManagers targeted itself, and picks the last created object to process. |
| `attachedLaunchObjects` _[AttachedLaunchObject](#attachedlaunchobject) array_ | Attached launch object information. A LaunchManager can be attached with a label on it with key `robolaunch.io/target-robot` and value of the target robot's name. Multiple LaunchManager could work together if they targeted the same Robot. |
| `initialBuildManagerStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | [*alpha*] Initial build manager creation status if exists. |
| `initialLaunchManagerStatuses` _[OwnedResourceStatus](#ownedresourcestatus) array_ | [*alpha*] Initial launch manager creation status if exists. |
| `attachedDevObjects` _[AttachedDevObject](#attacheddevobject) array_ | [*alpha*] Attached dev object information. |


#### WorkspaceManager



WorkspaceManager configures the ROS 2 workspaces and repositories by executing Kubernetes jobs.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `robot.roboscale.io/v1alpha1`
| `kind` _string_ | `WorkspaceManager`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[WorkspaceManagerSpec](#workspacemanagerspec)_ | Specification of the desired behavior of the WorkspaceManager. |
| `status` _[WorkspaceManagerStatus](#workspacemanagerstatus)_ | Most recently observed status of the WorkspaceManager. |


#### WorkspaceManagerSpec



WorkspaceManagerSpec defines the desired state of WorkspaceManager.

_Appears in:_
- [RobotSpec](#robotspec)
- [WorkspaceManager](#workspacemanager)

| Field | Description |
| --- | --- |
| `workspacesPath` _string_ | Global path of workspaces. It's fixed to `/root/workspaces` path. |
| `workspaces` _[Workspace](#workspace) array_ | Workspace definitions of robot. Multiple ROS 2 workspaces can be configured over this field. |
| `updateNeeded` _boolean_ | WorkspaceManager is triggered if this field is set to `true`. Then the workspaces are being configured again while backing up the old configurations. This field is often used by operator. |


#### WorkspaceManagerStatus



WorkspaceManagerStatus defines the observed state of WorkspaceManager.

_Appears in:_
- [WorkspaceManager](#workspacemanager)
- [WorkspaceManagerInstanceStatus](#workspacemanagerinstancestatus)

| Field | Description |
| --- | --- |
| `phase` _[WorkspaceManagerPhase](#workspacemanagerphase)_ | Phase of WorkspaceManager. |
| `clonerJobStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of cloner job. |
| `cleanupJobStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of cleanup jobs that runs while reconfiguring workspaces. |
| `version` _integer_ | Incremental version of workspace configuration map. Used to determine changes in configuration. |


#### BuildManager



BuildManager is the Schema for the buildmanagers API



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `robot.roboscale.io/v1alpha1`
| `kind` _string_ | `BuildManager`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[BuildManagerSpec](#buildmanagerspec)_ | Specification of the desired behavior of the BuildManager. |
| `status` _[BuildManagerStatus](#buildmanagerstatus)_ | Most recently observed status of the BuildManager. |


#### BuildManagerSpec



BuildManagerSpec defines the desired state of BuildManager.

_Appears in:_
- [BuildManager](#buildmanager)
- [RobotSpec](#robotspec)

| Field | Description |
| --- | --- |
| `steps` _[Step](#step) array_ | Defines the building steps. |


#### BuildManagerStatus



BuildManagerStatus defines the observed state of BuildManager.

_Appears in:_
- [AttachedBuildObject](#attachedbuildobject)
- [BuildManager](#buildmanager)

| Field | Description |
| --- | --- |
| `phase` _[BuildManagerPhase](#buildmanagerphase)_ | Phase of BuildManager. |
| `active` _boolean_ | Indicates if the BuildManager is currently executing it's jobs. |
| `scriptConfigMapStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of the ConfigMap that holds scripts. If a script is specified inside `.spec.steps[k]`, they are mounted to the step jobs via this ConfigMap. |
| `steps` _[StepStatus](#stepstatus) array_ | Statuses of the build steps. |


#### LaunchManager



LaunchManager is the Schema for the launchmanagers API



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `robot.roboscale.io/v1alpha1`
| `kind` _string_ | `LaunchManager`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[LaunchManagerSpec](#launchmanagerspec)_ | Specification of the desired behavior of the LaunchManager. |
| `status` _[LaunchManagerStatus](#launchmanagerstatus)_ | Most recently observed status of the LaunchManager. |


#### LaunchManagerSpec



LaunchManagerSpec defines the desired state of LaunchManager.

_Appears in:_
- [LaunchManager](#launchmanager)
- [RobotSpec](#robotspec)

| Field | Description |
| --- | --- |
| `launches` _object (keys:string, values:[Launch](#launch))_ | Launch descriptions. Every object defined here generates a launching command in the specified workspace. |


#### LaunchManagerStatus



LaunchManagerStatus defines the observed state of LaunchManager.

_Appears in:_
- [AttachedLaunchObject](#attachedlaunchobject)
- [LaunchManager](#launchmanager)

| Field | Description |
| --- | --- |
| `phase` _[LaunchManagerPhase](#launchmanagerphase)_ | Phase of LaunchManager. |
| `active` _boolean_ | Indicates if the LaunchManager is attached to a Robot and actively running. |
| `launchPodStatus` _[LaunchPodStatus](#launchpodstatus)_ | Collective statuses of launch pod and launch objects. |


#### RobotDevSuite



RobotDevSuite is a custom resource that creates dynamically configured development environments for robots.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `robot.roboscale.io/v1alpha1`
| `kind` _string_ | `RobotDevSuite`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[RobotDevSuiteSpec](#robotdevsuitespec)_ | Specification of the desired behavior of the RobotDevSuite. |
| `status` _[RobotDevSuiteStatus](#robotdevsuitestatus)_ | Most recently observed status of the RobotDevSuite. |


#### RobotDevSuiteSpec



RobotDevSuiteSpec defines the desired state of RobotDevSuite.

_Appears in:_
- [RobotDevSuite](#robotdevsuite)
- [RobotSpec](#robotspec)

| Field | Description |
| --- | --- |
| `vdiEnabled` _boolean_ | If `true`, a Cloud VDI will be provisioned inside development suite. |
| `robotVDITemplate` _[RobotVDISpec](#robotvdispec)_ | Configurational parameters of RobotVDI. Applied if `.spec.vdiEnabled` is set to `true`. |
| `ideEnabled` _boolean_ | If `true`, a Cloud IDE will be provisioned inside development suite. |
| `robotIDETemplate` _[RobotIDESpec](#robotidespec)_ | Configurational parameters of RobotIDE. Applied if `.spec.ideEnabled` is set to `true`. |


#### RobotDevSuiteStatus



RobotDevSuiteStatus defines the observed state of RobotDevSuite.

_Appears in:_
- [AttachedDevObject](#attacheddevobject)
- [RobotDevSuite](#robotdevsuite)
- [RobotDevSuiteInstanceStatus](#robotdevsuiteinstancestatus)

| Field | Description |
| --- | --- |
| `phase` _[RobotDevSuitePhase](#robotdevsuitephase)_ | Phase of RobotDevSuite. |
| `robotVDIStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of RobotVDI. |
| `robotIDEStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of RobotIDE. |
| `active` _boolean_ | [*alpha*] Indicates if RobotDevSuite is attached to a Robot and actively provisioned it's resources. |


#### MetricsExporter



MetricsExporter collects metrics from host machine and expose them from the Kubernetes API.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `robot.roboscale.io/v1alpha1`
| `kind` _string_ | `MetricsExporter`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[MetricsExporterSpec](#metricsexporterspec)_ | Specification of the desired behavior of the MetricsExporter. |
| `status` _[MetricsExporterStatus](#metricsexporterstatus)_ | Most recently observed status of the MetricsExporter. |


#### MetricsExporterSpec



MetricsExporterSpec defines the desired state of MetricsExporter.

_Appears in:_
- [MetricsExporter](#metricsexporter)

| Field | Description |
| --- | --- |
| `gpu` _[GPUMetrics](#gpumetrics)_ | Configurational parameters about GPU metrics collection. |
| `network` _[NetworkMetrics](#networkmetrics)_ | Configurational parameters about network metrics collection. |


#### MetricsExporterStatus



MetricsExporterStatus defines the observed state of MetricsExporter.

_Appears in:_
- [MetricsExporter](#metricsexporter)

| Field | Description |
| --- | --- |
| `phase` _[MetricsExporterPhase](#metricsexporterphase)_ | Phase of MetricsExporter. |
| `roleStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of role created for main and sidecar applications. |
| `roleBindingStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of role binding created for main and sidecar applications. |
| `saStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of service account created for main and sidecar applications. |
| `podStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of MetricsExporter pod. |
| `usage` _[Usage](#usage)_ | Usage metrics. |


#### RobotVDI



RobotVDI creates and manages Cloud VDI resources and workloads.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `robot.roboscale.io/v1alpha1`
| `kind` _string_ | `RobotVDI`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[RobotVDISpec](#robotvdispec)_ | Specification of the desired behavior of the RobotVDI. |
| `status` _[RobotVDIStatus](#robotvdistatus)_ | Most recently observed status of the RobotVDI. |


#### RobotVDISpec



RobotVDISpec defines the desired state of RobotVDI.

_Appears in:_
- [RobotDevSuiteSpec](#robotdevsuitespec)
- [RobotVDI](#robotvdi)

| Field | Description |
| --- | --- |
| `resources` _[Resources](#resources)_ | Resource limitations of Cloud IDE. |
| `serviceType` _[ServiceType](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#servicetype-v1-core)_ | Service type of Cloud IDE. `ClusterIP` and `NodePort` is supported. |
| `privileged` _boolean_ | If `true`, containers of RobotIDE will be privileged containers. It can be used in physical instances where it's necessary to access I/O devices on the host machine. Not recommended to activate this field on cloud instances. |
| `nat1to1` _string_ | NAT1TO1 option for Cloud VDI. |
| `webrtcPortRange` _string_ | UDP port range to used in WebRTC connections. |
| `resolution` _string_ | VDI screen resolution options. Default is `2048x1152`. |
| `ingress` _boolean_ | [*alpha*] RobotIDE will create an Ingress resource if `true`. |


#### RobotVDIStatus



RobotVDIStatus defines the observed state of RobotVDI.

_Appears in:_
- [RobotVDI](#robotvdi)

| Field | Description |
| --- | --- |
| `phase` _[RobotVDIPhase](#robotvdiphase)_ | Phase of RobotVDI. |
| `podStatus` _[OwnedPodStatus](#ownedpodstatus)_ | Status of Cloud VDI pod. |
| `serviceTCPStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of Cloud VDI TCP service. |
| `serviceUDPStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of Cloud VDI UDP service. |
| `ingressStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of Cloud VDI Ingress. |
| `pvcStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of Cloud VDI persistent volume claim. This PVC dynamically provisions a volume that is a shared between RobotVDI workloads and other workloads that requests display. |


#### RobotIDE



RobotIDE creates and manages Cloud IDE resources and workloads.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `robot.roboscale.io/v1alpha1`
| `kind` _string_ | `RobotIDE`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[RobotIDESpec](#robotidespec)_ | Specification of the desired behavior of the RobotIDE. |
| `status` _[RobotIDEStatus](#robotidestatus)_ | Most recently observed status of the RobotIDE. |


#### RobotIDESpec



RobotIDESpec defines the desired state of RobotIDE.

_Appears in:_
- [RobotDevSuiteSpec](#robotdevsuitespec)
- [RobotIDE](#robotide)

| Field | Description |
| --- | --- |
| `resources` _[Resources](#resources)_ | Resource limitations of Cloud IDE. |
| `serviceType` _[ServiceType](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#servicetype-v1-core)_ | Service type of Cloud IDE. `ClusterIP` and `NodePort` is supported. |
| `privileged` _boolean_ | If `true`, containers of RobotIDE will be privileged containers. It can be used in physical instances where it's necessary to access I/O devices on the host machine. Not recommended to activate this field on cloud instances. |
| `display` _boolean_ | Cloud IDE connects an X11 socket if it's set to `true` and a target RobotVDI resource is set in labels with key `robolaunch.io/target-vdi`. Applications that requires GUI can be executed such as rViz. |
| `ingress` _boolean_ | [*alpha*] RobotIDE will create an Ingress resource if `true`. |


#### RobotIDEStatus



RobotIDEStatus defines the observed state of RobotIDE.

_Appears in:_
- [RobotIDE](#robotide)

| Field | Description |
| --- | --- |
| `phase` _[RobotIDEPhase](#robotidephase)_ | Phase of RobotIDE. |
| `podStatus` _[OwnedPodStatus](#ownedpodstatus)_ | Status of Cloud IDE pod. |
| `serviceStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of Cloud IDE service. |
| `ingressStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of Cloud IDE Ingress. |


#### DiscoveryServer



DiscoveryServer is a custom resource that connects Robots and Fleets both locally and geoghraphically in DDS (UDP multicast) level.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `robot.roboscale.io/v1alpha1`
| `kind` _string_ | `DiscoveryServer`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[DiscoveryServerSpec](#discoveryserverspec)_ | Specification of the desired behavior of the DiscoveryServer. |
| `status` _[DiscoveryServerStatus](#discoveryserverstatus)_ | Most recently observed status of the DiscoveryServer. |


#### DiscoveryServerSpec



DiscoveryServerSpec defines the desired state of DiscoveryServer.

_Appears in:_
- [DiscoveryServer](#discoveryserver)
- [RobotSpec](#robotspec)

| Field | Description |
| --- | --- |
| `type` _[DiscoveryServerInstanceType](#discoveryserverinstancetype)_ | Instance type can be either `Server` or `Client`. If `Server`, instance creates discovery server resources and workloads. Other `Client` instances can connect to the `Server` instance. If `Client`, instance tries to connect a `Server` instance and hold `Server` configuration in a ConfigMap. |
| `reference` _[ObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectreference-v1-core)_ | Reference to the `Server` instance. It is used if `.spec.type` is `Client`. Referenced object can be previously provisioned in another cluster. In that case, cluster's name can be specified in `.spec.cluster` field. |
| `cluster` _string_ | Cloud instance name that holds DiscoveryServer instance with `Server` type. Should be empty if the type is `Server` since it takes cloud instance's name automatically. Should be set if the type is `Client`. |
| `hostname` _string_ | If instance type is `Server`, it can be an arbitrary value. If instance type is `Client`, it should be the same with Server's hostname. Used for getting Server's IP over DNS. |
| `subdomain` _string_ | If instance type is `Server`, it can be an arbitrary value. If instance type is `Client`, it should be the same with Server's subdomain. Used for getting Server's IP over DNS. |


#### DiscoveryServerStatus



DiscoveryServerStatus defines the observed state of DiscoveryServer.

_Appears in:_
- [DiscoveryServer](#discoveryserver)
- [DiscoveryServerInstanceStatus](#discoveryserverinstancestatus)

| Field | Description |
| --- | --- |
| `phase` _DiscoveryServerPhase_ | Phase of the DiscoveryServer. |
| `serviceStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of the DiscoveryServer service. |
| `serviceExportStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of the DiscoveryServer ServiceExport. |
| `podStatus` _[OwnedPodStatus](#ownedpodstatus)_ | Status of the DiscoveryServer pod. |
| `configMapStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of the DiscoveryServer config map. |
| `connectionInfo` _[ConnectionInfo](#connectioninfo)_ | Connection information. |


#### ROSBridge



ROSBridge is a custom resource that provisions ROS/2 bridge resources and workloads. It could also convert ROS 2 topics to ROS topics using ROS 1 to 2 bridge. (see https://github.com/ros2/ros1_bridge)



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `robot.roboscale.io/v1alpha1`
| `kind` _string_ | `ROSBridge`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `spec` _[ROSBridgeSpec](#rosbridgespec)_ | Specification of the desired behavior of the ROSBridge. |
| `status` _[ROSBridgeStatus](#rosbridgestatus)_ | Most recently observed status of the ROSBridge. |


#### ROSBridgeSpec



ROSBridgeSpec defines the desired state of ROSBridge.

_Appears in:_
- [ROSBridge](#rosbridge)
- [RobotSpec](#robotspec)

| Field | Description |
| --- | --- |
| `ros` _[BridgeDistro](#bridgedistro)_ | Configurational parameters for ROS bridge. |
| `ros2` _[BridgeDistro](#bridgedistro)_ | Configurational parameters for ROS 2 bridge. |


#### ROSBridgeStatus



ROSBridgeStatus defines the observed state of ROSBridge.

_Appears in:_
- [ROSBridge](#rosbridge)
- [ROSBridgeInstanceStatus](#rosbridgeinstancestatus)

| Field | Description |
| --- | --- |
| `phase` _BridgePhase_ | Phase of ROSBridge. |
| `podStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of ROSBridge pod. |
| `serviceStatus` _[OwnedResourceStatus](#ownedresourcestatus)_ | Status of ROSBridge service. |


#### RobotArtifact



RobotArtifact is a non-functional resource that holds Robot's specifications. It is used to define Robot templates.



| Field | Description |
| --- | --- |
| `apiVersion` _string_ | `robot.roboscale.io/v1alpha1`
| `kind` _string_ | `RobotArtifact`
| `metadata` _[ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectmeta-v1-meta)_ | Refer to Kubernetes API documentation for fields of `metadata`. |
| `template` _[RobotSpec](#robotspec)_ | Holds Robot's `.spec`. |


#### AttachedBuildObject





_Appears in:_
- [RobotStatus](#robotstatus)

| Field | Description |
| --- | --- |
| `reference` _[ObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectreference-v1-core)_ | Reference to the BuildManager instance. |
| `status` _[BuildManagerStatus](#buildmanagerstatus)_ | Status of attached BuildManager. |


#### AttachedDevObject





_Appears in:_
- [RobotStatus](#robotstatus)

| Field | Description |
| --- | --- |
| `reference` _[ObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectreference-v1-core)_ | Reference to the RobotDevSuite instance. |
| `status` _[RobotDevSuiteStatus](#robotdevsuitestatus)_ | Status of attached RobotDevSuite. |


#### AttachedLaunchObject





_Appears in:_
- [RobotStatus](#robotstatus)

| Field | Description |
| --- | --- |
| `reference` _[ObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectreference-v1-core)_ | Reference to the LaunchManager instance. |
| `status` _[LaunchManagerStatus](#launchmanagerstatus)_ | Status of attached LaunchManager. |


#### BridgeDistro





_Appears in:_
- [ROSBridgeSpec](#rosbridgespec)

| Field | Description |
| --- | --- |
| `enabled` _boolean_ | If `true`, resources and workloads are created by ROSBridge. |
| `distro` _[ROSDistro](#rosdistro)_ | ROS distribution for bridge. |


#### BuildManagerPhase

_Underlying type:_ `string`



_Appears in:_
- [BuildManagerStatus](#buildmanagerstatus)



#### ConnectionInfo





_Appears in:_
- [DiscoveryServerStatus](#discoveryserverstatus)

| Field | Description |
| --- | --- |
| `uri` _string_ | URI of the discovery server. Discovery server instance tries to ping this address to see if it's reachable. |
| `ip` _string_ | IP of the discovery server. IP is being obtained from the DNS name, which is being built according to the discovery server configuration. |
| `configMapName` _string_ | Name of the config map that holds discovery server configuration. |


#### DiscoveryServerInstanceStatus





_Appears in:_
- [RobotStatus](#robotstatus)

| Field | Description |
| --- | --- |
| `resource` _[OwnedResourceStatus](#ownedresourcestatus)_ | Generic status for any owned resource. |
| `status` _[DiscoveryServerStatus](#discoveryserverstatus)_ | Status of the DiscoveryServer instance. |


#### DiscoveryServerInstanceType

_Underlying type:_ `string`

Instance type can be either `Server` or `Client`.

_Appears in:_
- [DiscoveryServerSpec](#discoveryserverspec)



#### GPUMetrics





_Appears in:_
- [MetricsExporterSpec](#metricsexporterspec)

| Field | Description |
| --- | --- |
| `track` _boolean_ | MetricsExporter watches volatile GPU usage in the host machine if it's set to `true`. |
| `interval` _integer_ | Watching latency. |


#### GPUUtilizationStatus





_Appears in:_
- [Usage](#usage)

| Field | Description |
| --- | --- |
| `utilization` _string_ | Volatile GPU utilization. Shows a percentage gathered from `nvidia-smi` command. |
| `lastUpdateTimestamp` _string_ | Last update time. |


#### Launch



Launch description of a repository.

_Appears in:_
- [LaunchManagerSpec](#launchmanagerspec)

| Field | Description |
| --- | --- |
| `instances` _string array_ | Cluster selector. If the current instance name is on the list, LaunchManager creates launch pods. |
| `workspace` _string_ | Name of the workspace. Should be selected among the existing workspaces in WorkspaceManager's manifests. |
| `namespacing` _boolean_ | ROS 2 namespacing. May not be suitable for all launchfiles. If used, all the node names and topic names should be defined relative, not absolute. (eg. `cmd_vel` instead of /cmd_vel``) |
| `entrypoint` _[LaunchEntrypointConfig](#launchentrypointconfig)_ | Entrypoint configuration of launch. |
| `container` _[LaunchContainerConfig](#launchcontainerconfig)_ | General container configuration parameters. |


#### LaunchContainerConfig





_Appears in:_
- [Launch](#launch)

| Field | Description |
| --- | --- |
| `env` _[EnvVar](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#envvar-v1-core) array_ | Additional environment variables to set when launching ROS nodes. |
| `privileged` _boolean_ | Launch container privilege. |
| `resources` _[Resources](#resources)_ | Launch container resource limits. |
| `display` _boolean_ | Launch processes connects an X11 socket if it's set to `true` and a target RobotVDI resource is set in labels with key `robolaunch.io/target-vdi`. Applications that requires GUI can be executed such as rViz. |


#### LaunchEntrypointConfig





_Appears in:_
- [Launch](#launch)

| Field | Description |
| --- | --- |
| `type` _[LaunchType](#launchtype)_ | Launching type. Can be `Launch`, `Run` or `Custom`. |
| `package` _string_ | Package name. (eg. `robolaunch_cloudy_navigation`) |
| `launchfile` _string_ | Launchfile. (eg. `nav_launch.py`) Required and used if the launch type is `Launch`. |
| `executable` _string_ | Executable file name. (eg. `webcam_pub.py`) Required and used if the launch type is `Run`. |
| `disableSourcingWs` _boolean_ | If `true`, workspaces are not sourced by default. Used if the launch type is `Custom`. |
| `cmd` _string_ | Custom command to launch packages or start nodes. Required if the launch type is `Custom`. |
| `parameters` _object (keys:string, values:string)_ | Launch parameters. |
| `prelaunch` _[Prelaunch](#prelaunch)_ | Command or script to run just before node's execution. |


#### LaunchManagerPhase

_Underlying type:_ `string`



_Appears in:_
- [LaunchManagerStatus](#launchmanagerstatus)



#### LaunchPodStatus





_Appears in:_
- [LaunchManagerStatus](#launchmanagerstatus)

| Field | Description |
| --- | --- |
| `status` _[OwnedPodStatus](#ownedpodstatus)_ | Launch pod status. Every LaunchManager creates one pod if active. |
| `launchStatus` _object (keys:string, values:[LaunchStatus](#launchstatus))_ | Status of launch objects. Every launch object generates a `ros2 launch` command that will run as an entrypoint in a container. |


#### LaunchStatus





_Appears in:_
- [LaunchPodStatus](#launchpodstatus)

| Field | Description |
| --- | --- |
| `active` _boolean_ | Inditaces if the launch process are actively running on cluster. It may not be selected by launch cluster selectors. |
| `containerStatus` _[ContainerStatus](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#containerstatus-v1-core)_ | Statuses of the containers of pods owned by LaunchManager. |


#### LaunchType

_Underlying type:_ `string`



_Appears in:_
- [LaunchEntrypointConfig](#launchentrypointconfig)



#### MetricsExporterPhase

_Underlying type:_ `string`



_Appears in:_
- [MetricsExporterStatus](#metricsexporterstatus)



#### NetworkInterfaceLoad





_Appears in:_
- [NetworkLoadStatus](#networkloadstatus)

| Field | Description |
| --- | --- |
| `in` _string_ | Average load of incoming packets. |
| `out` _string_ | Average load of outgoing packets. |


#### NetworkLoadStatus





_Appears in:_
- [Usage](#usage)

| Field | Description |
| --- | --- |
| `load` _object (keys:string, values:[NetworkInterfaceLoad](#networkinterfaceload))_ | Loads values of network interfaces. |
| `lastUpdateTimestamp` _string_ | Last update time. |


#### NetworkMetrics





_Appears in:_
- [MetricsExporterSpec](#metricsexporterspec)

| Field | Description |
| --- | --- |
| `track` _boolean_ | MetricsExporter watches network loads in the host machine if it's set to `true`. |
| `interval` _integer_ | Watching latency. |
| `interfaces` _string array_ | Network interfaces which are desired to being watched. |


#### OwnedPodStatus





_Appears in:_
- [DiscoveryServerStatus](#discoveryserverstatus)
- [LaunchPodStatus](#launchpodstatus)
- [RobotIDEStatus](#robotidestatus)
- [RobotVDIStatus](#robotvdistatus)

| Field | Description |
| --- | --- |
| `resource` _[OwnedResourceStatus](#ownedresourcestatus)_ | Generic status for any owned resource. |
| `ip` _string_ | IP of the pod. |


#### OwnedResourceStatus



Generic status for any owned resource.

_Appears in:_
- [BuildManagerStatus](#buildmanagerstatus)
- [DiscoveryServerInstanceStatus](#discoveryserverinstancestatus)
- [DiscoveryServerStatus](#discoveryserverstatus)
- [MetricsExporterStatus](#metricsexporterstatus)
- [OwnedPodStatus](#ownedpodstatus)
- [ROSBridgeInstanceStatus](#rosbridgeinstancestatus)
- [ROSBridgeStatus](#rosbridgestatus)
- [RobotDevSuiteInstanceStatus](#robotdevsuiteinstancestatus)
- [RobotDevSuiteStatus](#robotdevsuitestatus)
- [RobotIDEStatus](#robotidestatus)
- [RobotStatus](#robotstatus)
- [RobotVDIStatus](#robotvdistatus)
- [StepStatus](#stepstatus)
- [VolumeStatuses](#volumestatuses)
- [WorkspaceManagerInstanceStatus](#workspacemanagerinstancestatus)
- [WorkspaceManagerStatus](#workspacemanagerstatus)

| Field | Description |
| --- | --- |
| `created` _boolean_ | Shows if the owned resource is created. |
| `reference` _[ObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#objectreference-v1-core)_ | Reference to the owned resource. |
| `phase` _string_ | Phase of the owned resource. |


#### Prelaunch



Prelaunch command or script is applied just before the node is started.

_Appears in:_
- [LaunchEntrypointConfig](#launchentrypointconfig)

| Field | Description |
| --- | --- |
| `command` _string_ | Bash command to run before ROS node execution. |


#### RMWImplementation

_Underlying type:_ `string`

RMW implementation chooses DDS vendor for ROS 2. Currently, only eProsima's FastDDS is supported.

_Appears in:_
- [RobotSpec](#robotspec)



#### ROSBridgeInstanceStatus





_Appears in:_
- [RobotStatus](#robotstatus)

| Field | Description |
| --- | --- |
| `resource` _[OwnedResourceStatus](#ownedresourcestatus)_ | Generic status for any owned resource. |
| `status` _[ROSBridgeStatus](#rosbridgestatus)_ | Status of the ROSBridge instance. |


#### ROSDistro

_Underlying type:_ `string`

ROS 2 distribution selection. Currently supported distributions are Humble, Foxy, Galactic.

_Appears in:_
- [BridgeDistro](#bridgedistro)
- [RobotSpec](#robotspec)
- [Workspace](#workspace)



#### Repository



Repository description.

_Appears in:_
- [Workspace](#workspace)

| Field | Description |
| --- | --- |
| `url` _string_ | Base URL of the repository. |
| `branch` _string_ | Branch of the repository to clone. |
| `path` _string_ | [*Autofilled*] Absolute path of repository |
| `owner` _string_ | [*Autofilled*] User or organization, maintainer of repository |
| `repo` _string_ | [*Autofilled*] Repository name |
| `hash` _string_ | [*Autofilled*] Hash of last commit |


#### Resources



VDI resource limits.

_Appears in:_
- [LaunchContainerConfig](#launchcontainerconfig)
- [RobotIDESpec](#robotidespec)
- [RobotVDISpec](#robotvdispec)

| Field | Description |
| --- | --- |
| `gpuCore` _integer_ | GPU core number that will be allocated. |
| `cpu` _string_ | CPU resource limit. |
| `memory` _string_ | Memory resource limit. |


#### RobotDevSuiteInstanceStatus





_Appears in:_
- [RobotStatus](#robotstatus)

| Field | Description |
| --- | --- |
| `resource` _[OwnedResourceStatus](#ownedresourcestatus)_ | Generic status for any owned resource. |
| `status` _[RobotDevSuiteStatus](#robotdevsuitestatus)_ | Status of the RobotDevSuite instance. |


#### RobotDevSuitePhase

_Underlying type:_ `string`



_Appears in:_
- [RobotDevSuiteStatus](#robotdevsuitestatus)



#### RobotIDEPhase

_Underlying type:_ `string`



_Appears in:_
- [RobotIDEStatus](#robotidestatus)



#### RobotVDIPhase

_Underlying type:_ `string`



_Appears in:_
- [RobotVDIStatus](#robotvdistatus)



#### RootDNSConfig





_Appears in:_
- [RobotSpec](#robotspec)

| Field | Description |
| --- | --- |
| `host` _string_ | [*alpha*] Root DNS name.. |


#### Step



Step is a command or script to execute when building a robot. Either `command` or `script` should be specified for each step.

_Appears in:_
- [BuildManagerSpec](#buildmanagerspec)
- [StepStatus](#stepstatus)

| Field | Description |
| --- | --- |
| `instances` _string array_ | Cluster selector. If the current instance name is on the list, BuildManager creates building jobs. |
| `name` _string_ | Name of the step. |
| `workspace` _string_ | Name of the workspace. Should be selected among the existing workspaces in WorkspaceManager's manifests. |
| `command` _string_ | Bash command to run. Assume that your command will be `/bin/bash -c <COMMAND>`. Use logical operators (eg. `&&`) and pipes if the multiple dependent commands will be executed. |
| `script` _string_ | Bash script to run. |
| `env` _[EnvVar](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#envvar-v1-core) array_ | Environment variables for step. |


#### StepStatus





_Appears in:_
- [BuildManagerStatus](#buildmanagerstatus)

| Field | Description |
| --- | --- |
| `resource` _[OwnedResourceStatus](#ownedresourcestatus)_ | Generic status for any owned resource. |
| `step` _[Step](#step)_ | Status of the step. |


#### Storage



Robot's resource limitations.

_Appears in:_
- [RobotSpec](#robotspec)

| Field | Description |
| --- | --- |
| `amount` _integer_ | Specifies how much storage will be allocated in total. Use MB as a unit of measurement. (eg. `10240` is equal to 10 GB) |
| `storageClassConfig` _[StorageClassConfig](#storageclassconfig)_ | Storage class selection for robot's volumes. |


#### StorageClassConfig



Storage class configuration for a volume type.

_Appears in:_
- [Storage](#storage)

| Field | Description |
| --- | --- |
| `name` _string_ | Storage class name. |
| `accessMode` _[PersistentVolumeAccessMode](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.22/#persistentvolumeaccessmode-v1-core)_ | PVC access modes. Currently, only `ReadWriteOnce` is supported. |


#### TLSSecretReference





_Appears in:_
- [RobotSpec](#robotspec)

| Field | Description |
| --- | --- |
| `name` _string_ | [*alpha*] TLS secret object name. |
| `namespace` _string_ | [*alpha*] TLS secret object namespace. |


#### Usage





_Appears in:_
- [MetricsExporterStatus](#metricsexporterstatus)

| Field | Description |
| --- | --- |
| `gpu` _[GPUUtilizationStatus](#gpuutilizationstatus)_ | GPU usage information. |
| `network` _[NetworkLoadStatus](#networkloadstatus)_ | Network usage information. |


#### VolumeStatuses





_Appears in:_
- [RobotStatus](#robotstatus)

| Field | Description |
| --- | --- |
| `varDir` _[OwnedResourceStatus](#ownedresourcestatus)_ | Holds PVC status of the `/var` directory of underlying OS. |
| `etcDir` _[OwnedResourceStatus](#ownedresourcestatus)_ | Holds PVC status of the `/etc` directory of underlying OS. |
| `usrDir` _[OwnedResourceStatus](#ownedresourcestatus)_ | Holds PVC status of the `/usr` directory of underlying OS. |
| `optDir` _[OwnedResourceStatus](#ownedresourcestatus)_ | Holds PVC status of the `/opt` directory of underlying OS. |
| `workspaceDir` _[OwnedResourceStatus](#ownedresourcestatus)_ | Holds PVC status of the workspaces directory of underlying OS. |


#### Workspace



Workspace description. Each robot should contain at least one workspace. A workspace should contain at least one repository in it.

_Appears in:_
- [WorkspaceManagerSpec](#workspacemanagerspec)

| Field | Description |
| --- | --- |
| `name` _string_ | Name of workspace. If a workspace's name is `my_ws`, it's absolute path is `/home/workspaces/my_ws`. |
| `distro` _[ROSDistro](#rosdistro)_ |  |
| `repositories` _object (keys:string, values:[Repository](#repository))_ | Repositories to clone inside workspace's `src` directory. |


#### WorkspaceManagerInstanceStatus





_Appears in:_
- [RobotStatus](#robotstatus)

| Field | Description |
| --- | --- |
| `resource` _[OwnedResourceStatus](#ownedresourcestatus)_ | Generic status for any owned resource. |
| `status` _[WorkspaceManagerStatus](#workspacemanagerstatus)_ | Status of the WorkspaceManager instance. |


#### WorkspaceManagerPhase

_Underlying type:_ `string`



_Appears in:_
- [WorkspaceManagerStatus](#workspacemanagerstatus)


