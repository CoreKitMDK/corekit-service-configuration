helm install internal-configuration-kvstore bitnami/keydb --namespace testing-dev -f values.yaml


#Tin@X CLANGARM64 /c/repos/CoreKitMDK/corekit-service-configuration/deploy (main)
#$ helm uninstall internal-configuration-kvstore -n testing-dev
#release "internal-configuration-kvstore" uninstalled
#
#Tin@X CLANGARM64 /c/repos/CoreKitMDK/corekit-service-configuration/deploy (main)
#$ kubectl delete pvc -l app.kubernetes.io/instance=internal-configuration-kvstore -n testing-dev
#persistentvolumeclaim "data-internal-configuration-kvstore-master-0" deleted
