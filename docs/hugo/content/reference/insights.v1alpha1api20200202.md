---
---
<h2 id="insights.azure.com/v1alpha1api20200202">insights.azure.com/v1alpha1api20200202</h2>
<div>
<p>Package v1alpha1api20200202 contains API Schema definitions for the insights v1alpha1api20200202 API group</p>
</div>
Resource Types:
<ul></ul>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.Components_SpecARM">Components_SpecARM</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponentProperties. Use v1beta20200202.ApplicationInsightsComponentProperties instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Application_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesApplicationType">
ApplicationInsightsComponentPropertiesApplicationType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>DisableIpMasking</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>DisableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Flow_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesFlowType">
ApplicationInsightsComponentPropertiesFlowType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ForceCustomerStorageForProfiler</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>HockeyAppId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ImmediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>IngestionMode</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesIngestionMode">
ApplicationInsightsComponentPropertiesIngestionMode
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion">
ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery">
ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Request_Source</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesRequestSource">
ApplicationInsightsComponentPropertiesRequestSource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>RetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>SamplingPercentage</code><br/>
<em>
float64
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>workspaceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesApplicationType">ApplicationInsightsComponentPropertiesApplicationType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponentPropertiesApplicationType. Use
v1beta20200202.ApplicationInsightsComponentPropertiesApplicationType instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;other&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;web&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesFlowType">ApplicationInsightsComponentPropertiesFlowType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponentPropertiesFlowType. Use
v1beta20200202.ApplicationInsightsComponentPropertiesFlowType instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Bluefield&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesIngestionMode">ApplicationInsightsComponentPropertiesIngestionMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponentPropertiesIngestionMode. Use
v1beta20200202.ApplicationInsightsComponentPropertiesIngestionMode instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;ApplicationInsights&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ApplicationInsightsWithDiagnosticSettings&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LogAnalytics&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion">ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion. Use
v1beta20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery">ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery. Use
v1beta20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesRequestSource">ApplicationInsightsComponentPropertiesRequestSource
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">ApplicationInsightsComponentPropertiesARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponentPropertiesRequestSource. Use
v1beta20200202.ApplicationInsightsComponentPropertiesRequestSource instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;rest&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusApplicationType">ApplicationInsightsComponentPropertiesStatusApplicationType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponentPropertiesStatusApplicationType. Use
v1beta20200202.ApplicationInsightsComponentPropertiesStatusApplicationType instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;other&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;web&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusFlowType">ApplicationInsightsComponentPropertiesStatusFlowType
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponentPropertiesStatusFlowType. Use
v1beta20200202.ApplicationInsightsComponentPropertiesStatusFlowType instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Bluefield&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusIngestionMode">ApplicationInsightsComponentPropertiesStatusIngestionMode
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponentPropertiesStatusIngestionMode. Use
v1beta20200202.ApplicationInsightsComponentPropertiesStatusIngestionMode instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;ApplicationInsights&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;ApplicationInsightsWithDiagnosticSettings&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;LogAnalytics&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusRequestSource">ApplicationInsightsComponentPropertiesStatusRequestSource
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponentPropertiesStatusRequestSource. Use
v1beta20200202.ApplicationInsightsComponentPropertiesStatusRequestSource instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;rest&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_StatusARM">ApplicationInsightsComponent_StatusARM</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponentProperties_Status. Use v1beta20200202.ApplicationInsightsComponentProperties_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>AppId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ApplicationId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Application_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusApplicationType">
ApplicationInsightsComponentPropertiesStatusApplicationType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ConnectionString</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>CreationDate</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>DisableIpMasking</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>DisableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Flow_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusFlowType">
ApplicationInsightsComponentPropertiesStatusFlowType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ForceCustomerStorageForProfiler</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>HockeyAppId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>HockeyAppToken</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ImmediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>IngestionMode</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusIngestionMode">
ApplicationInsightsComponentPropertiesStatusIngestionMode
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>InstrumentationKey</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>LaMigrationDate</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>PrivateLinkScopedResources</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.PrivateLinkScopedResource_StatusARM">
[]PrivateLinkScopedResource_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Request_Source</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusRequestSource">
ApplicationInsightsComponentPropertiesStatusRequestSource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>RetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>SamplingPercentage</code><br/>
<em>
float64
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>TenantId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>WorkspaceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.Component">Component</a>)
</p>
<div>
<p>Deprecated version of ApplicationInsightsComponent_Status. Use v1beta20200202.ApplicationInsightsComponent_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>AppId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ApplicationId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Application_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusApplicationType">
ApplicationInsightsComponentPropertiesStatusApplicationType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>conditions</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#Condition">
[]genruntime/conditions.Condition
</a>
</em>
</td>
<td>
<p>Conditions: The observed state of the resource</p>
</td>
</tr>
<tr>
<td>
<code>ConnectionString</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>CreationDate</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>DisableIpMasking</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>DisableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Flow_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusFlowType">
ApplicationInsightsComponentPropertiesStatusFlowType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ForceCustomerStorageForProfiler</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>HockeyAppId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>HockeyAppToken</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ImmediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>IngestionMode</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusIngestionMode">
ApplicationInsightsComponentPropertiesStatusIngestionMode
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>InstrumentationKey</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>LaMigrationDate</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>PrivateLinkScopedResources</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.PrivateLinkScopedResource_Status">
[]PrivateLinkScopedResource_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties_name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>provisioningState</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.PublicNetworkAccessType_Status">
PublicNetworkAccessType_Status
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Request_Source</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesStatusRequestSource">
ApplicationInsightsComponentPropertiesStatusRequestSource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>RetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>SamplingPercentage</code><br/>
<em>
float64
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
Kubernetes v1.JSON
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>TenantId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>WorkspaceResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_StatusARM">ApplicationInsightsComponent_StatusARM
</h3>
<div>
<p>Deprecated version of ApplicationInsightsComponent_Status. Use v1beta20200202.ApplicationInsightsComponent_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>id</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">
ApplicationInsightsComponentProperties_StatusARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
<a href="https://pkg.go.dev/k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1#JSON">
Kubernetes v1.JSON
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>type</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.Component">Component
</h3>
<div>
<p>Deprecated version of Component. Use v1beta20200202.Component instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>metadata</code><br/>
<em>
<a href="https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta">
Kubernetes meta/v1.ObjectMeta
</a>
</em>
</td>
<td>
Refer to the Kubernetes API documentation for the fields of the
<code>metadata</code> field.
</td>
</tr>
<tr>
<td>
<code>spec</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.Components_Spec">
Components_Spec
</a>
</em>
</td>
<td>
<br/>
<br/>
<table>
<tr>
<td>
<code>Application_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesApplicationType">
ApplicationInsightsComponentPropertiesApplicationType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>DisableIpMasking</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>DisableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Flow_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesFlowType">
ApplicationInsightsComponentPropertiesFlowType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ForceCustomerStorageForProfiler</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>HockeyAppId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ImmediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>IngestionMode</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesIngestionMode">
ApplicationInsightsComponentPropertiesIngestionMode
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion">
ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery">
ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Request_Source</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesRequestSource">
ApplicationInsightsComponentPropertiesRequestSource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>RetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>SamplingPercentage</code><br/>
<em>
float64
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>workspaceResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
</td>
</tr>
</table>
</td>
</tr>
<tr>
<td>
<code>status</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">
ApplicationInsightsComponent_Status
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.Components_Spec">Components_Spec
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.Component">Component</a>)
</p>
<div>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>Application_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesApplicationType">
ApplicationInsightsComponentPropertiesApplicationType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>azureName</code><br/>
<em>
string
</em>
</td>
<td>
<p>AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
doesn&rsquo;t have to be.</p>
</td>
</tr>
<tr>
<td>
<code>DisableIpMasking</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>DisableLocalAuth</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Flow_Type</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesFlowType">
ApplicationInsightsComponentPropertiesFlowType
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ForceCustomerStorageForProfiler</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>HockeyAppId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ImmediatePurgeDataOn30Days</code><br/>
<em>
bool
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>IngestionMode</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesIngestionMode">
ApplicationInsightsComponentPropertiesIngestionMode
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>owner</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#KnownResourceReference">
genruntime.KnownResourceReference
</a>
</em>
</td>
<td>
<p>Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
reference to a resources.azure.com/ResourceGroup resource</p>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForIngestion</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion">
ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>publicNetworkAccessForQuery</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery">
ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>Request_Source</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesRequestSource">
ApplicationInsightsComponentPropertiesRequestSource
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>RetentionInDays</code><br/>
<em>
int
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>SamplingPercentage</code><br/>
<em>
float64
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>workspaceResourceReference</code><br/>
<em>
<a href="https://pkg.go.dev/github.com/Azure/azure-service-operator/v2/pkg/genruntime#ResourceReference">
genruntime.ResourceReference
</a>
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.Components_SpecARM">Components_SpecARM
</h3>
<div>
<p>Deprecated version of Components_Spec. Use v1beta20200202.Components_Spec instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>etag</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>kind</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>location</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>name</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>properties</code><br/>
<em>
<a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentPropertiesARM">
ApplicationInsightsComponentPropertiesARM
</a>
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>tags</code><br/>
<em>
map[string]string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.PrivateLinkScopedResource_Status">PrivateLinkScopedResource_Status
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status</a>)
</p>
<div>
<p>Deprecated version of PrivateLinkScopedResource_Status. Use v1beta20200202.PrivateLinkScopedResource_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ScopeId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.PrivateLinkScopedResource_StatusARM">PrivateLinkScopedResource_StatusARM
</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM</a>)
</p>
<div>
<p>Deprecated version of PrivateLinkScopedResource_Status. Use v1beta20200202.PrivateLinkScopedResource_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Field</th>
<th>Description</th>
</tr>
</thead>
<tbody>
<tr>
<td>
<code>ResourceId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
<tr>
<td>
<code>ScopeId</code><br/>
<em>
string
</em>
</td>
<td>
</td>
</tr>
</tbody>
</table>
<h3 id="insights.azure.com/v1alpha1api20200202.PublicNetworkAccessType_Status">PublicNetworkAccessType_Status
(<code>string</code> alias)</h3>
<p>
(<em>Appears on:</em><a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponentProperties_StatusARM">ApplicationInsightsComponentProperties_StatusARM</a>, <a href="#insights.azure.com/v1alpha1api20200202.ApplicationInsightsComponent_Status">ApplicationInsightsComponent_Status</a>)
</p>
<div>
<p>Deprecated version of PublicNetworkAccessType_Status. Use v1beta20200202.PublicNetworkAccessType_Status instead</p>
</div>
<table>
<thead>
<tr>
<th>Value</th>
<th>Description</th>
</tr>
</thead>
<tbody><tr><td><p>&#34;Disabled&#34;</p></td>
<td></td>
</tr><tr><td><p>&#34;Enabled&#34;</p></td>
<td></td>
</tr></tbody>
</table>
<hr/>