Limitations
===========

StorageClass reclaim policy
---------------------------

TopoLVM does not care about `Retain` [reclaim policy](https://kubernetes.io/docs/concepts/storage/storage-classes/#reclaim-policy)
because CSI volumes can be referenced only via PersistentVolumeClaims.

Ref: https://kubernetes.io/docs/concepts/storage/volumes/#csi

> The `csi` volume type does not support direct reference from Pod and may
> only be referenced in a Pod via a `PersistentVolumeClaim` object. 

Pod without PVC
---------------

TopoLVM expects that PVCs are created in advance of their Pods.
However, the TopoLVM webhook does not block the creation of a Pod when there are missing PVCs for the Pod.
This is because such usages are valid in other StorageClasses and the webhook cannot identify the StorageClasses without PVCs.
For such Pods, TopoLVM's extended scheduler will not work.

The typical usage of TopoLVM is using StatefulSet with volumeClaimTemplate.
