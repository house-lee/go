<map version="freeplane 1.5.9">
<!--To view this file, download free mind mapping software Freeplane from http://freeplane.sourceforge.net -->
<node TEXT="SoarGo" FOLDED="false" ID="ID_1279842490" CREATED="1493882651049" MODIFIED="1493883173329" STYLE="oval">
<font SIZE="18"/>
<hook NAME="MapStyle" zoom="0.75">
    <properties fit_to_viewport="false;"/>

<map_styles>
<stylenode LOCALIZED_TEXT="styles.root_node" STYLE="oval" UNIFORM_SHAPE="true" VGAP_QUANTITY="24.0 pt">
<font SIZE="24"/>
<stylenode LOCALIZED_TEXT="styles.predefined" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="default" COLOR="#000000" STYLE="fork">
<font NAME="SansSerif" SIZE="10" BOLD="false" ITALIC="false"/>
</stylenode>
<stylenode LOCALIZED_TEXT="defaultstyle.details"/>
<stylenode LOCALIZED_TEXT="defaultstyle.attributes">
<font SIZE="9"/>
</stylenode>
<stylenode LOCALIZED_TEXT="defaultstyle.note" COLOR="#000000" BACKGROUND_COLOR="#ffffff" TEXT_ALIGN="LEFT"/>
<stylenode LOCALIZED_TEXT="defaultstyle.floating">
<edge STYLE="hide_edge"/>
<cloud COLOR="#f0f0f0" SHAPE="ROUND_RECT"/>
</stylenode>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.user-defined" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="styles.topic" COLOR="#18898b" STYLE="fork">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.subtopic" COLOR="#cc3300" STYLE="fork">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.subsubtopic" COLOR="#669900">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.important">
<icon BUILTIN="yes"/>
</stylenode>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.AutomaticLayout" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="AutomaticLayout.level.root" COLOR="#000000" STYLE="oval" SHAPE_HORIZONTAL_MARGIN="10.0 pt" SHAPE_VERTICAL_MARGIN="10.0 pt">
<font SIZE="18"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,1" COLOR="#0033ff">
<font SIZE="16"/>
<edge COLOR="#ff0000"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,2" COLOR="#00b439">
<font SIZE="14"/>
<edge COLOR="#0000ff"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,3" COLOR="#990000">
<font SIZE="12"/>
<edge COLOR="#00ff00"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,4" COLOR="#111111">
<font SIZE="10"/>
<edge COLOR="#ff00ff"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,5">
<edge COLOR="#00ffff"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,6">
<edge COLOR="#7c0000"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,7">
<edge COLOR="#00007c"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,8">
<edge COLOR="#007c00"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,9">
<edge COLOR="#7c007c"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,10">
<edge COLOR="#007c7c"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,11">
<edge COLOR="#7c7c00"/>
</stylenode>
</stylenode>
</stylenode>
</map_styles>
</hook>
<hook NAME="AutomaticEdgeColor" COUNTER="2" RULE="ON_BRANCH_CREATION"/>
<node TEXT="interfaces" POSITION="right" ID="ID_576646748" CREATED="1493882665117" MODIFIED="1493882669638">
<edge COLOR="#ff0000"/>
<node TEXT="Coordinator" ID="ID_1615158151" CREATED="1493882670131" MODIFIED="1493882678441">
<node TEXT="Pulse( serverID ) error" ID="ID_1142783488" CREATED="1493882703607" MODIFIED="1493882715102">
<node TEXT="Redis: Set key with expiry" ID="ID_421905807" CREATED="1493883182361" MODIFIED="1493883192655"/>
<node TEXT="ZooKeeper: Create temporary znode" ID="ID_226969898" CREATED="1493883192815" MODIFIED="1493883211641"/>
<node TEXT="ETCD: Set key with ttl" ID="ID_198907102" CREATED="1493883212465" MODIFIED="1493883217782"/>
</node>
<node TEXT="IsServerAlive( serverID ) bool, error" ID="ID_1284551578" CREATED="1493882715414" MODIFIED="1493882741139">
<node TEXT="Redis: Check if key exists" ID="ID_10862314" CREATED="1493883219953" MODIFIED="1493883231605"/>
<node TEXT="ZooKeeper: Check if znode exists" ID="ID_1121031788" CREATED="1493883231798" MODIFIED="1493883239895"/>
<node TEXT="ETCD: Check if key exists" ID="ID_891155891" CREATED="1493883240055" MODIFIED="1493883249739"/>
</node>
<node TEXT="GenerateRequestID() requestID" ID="ID_116339191" CREATED="1493882725603" MODIFIED="1493882755520">
<node TEXT="Use UUID algorithm" ID="ID_333167859" CREATED="1493883251569" MODIFIED="1493883296294">
<node TEXT="serverID+internal_incremental_id" ID="ID_1076209860" CREATED="1493883324886" MODIFIED="1493883356364"/>
<node TEXT="UUID(serverID,timestamp)" ID="ID_1805531711" CREATED="1493883356916" MODIFIED="1493883376863"/>
</node>
<node TEXT="Use Coordination Cluster" ID="ID_1962395418" CREATED="1493883296574" MODIFIED="1493883322407">
<node TEXT="Redis: INCR requestID" ID="ID_187428091" CREATED="1493883379175" MODIFIED="1493883387242"/>
<node TEXT="ZooKeeper: Create new sequential znode" ID="ID_1106609614" CREATED="1493883388842" MODIFIED="1493883421087"/>
<node TEXT="ETCD: ?" ID="ID_1807321992" CREATED="1493883699081" MODIFIED="1493883703323"/>
</node>
</node>
<node TEXT="UpdateRequestProgress(requestID, serverID, workStationID, jobDesc)" ID="ID_13869073" CREATED="1493882755701" MODIFIED="1493919552252"/>
<node TEXT="GetRequestProgress( requestID) (serverID, workStationID)" ID="ID_1753990823" CREATED="1493882927528" MODIFIED="1493882947961"/>
<node TEXT="SaveRequestResult( requestID, header []byte, body []byte) error" ID="ID_995577680" CREATED="1493925545216" MODIFIED="1493925639298"/>
<node TEXT="GetRequestResult( requestID ) (header, body, error)" ID="ID_37602464" CREATED="1493925639497" MODIFIED="1493925668788"/>
<node TEXT="GetTaskList(workStationID) taskList" ID="ID_1919387610" CREATED="1493883744605" MODIFIED="1493883779039">
<node TEXT="A task contains:" ID="ID_755157473" CREATED="1493883912360" MODIFIED="1493883923854">
<node TEXT="serverID" ID="ID_676306034" CREATED="1493883930648" MODIFIED="1493883950229"/>
<node TEXT="requestID" ID="ID_287000149" CREATED="1493883953079" MODIFIED="1493884212439"/>
</node>
<node TEXT="Possible Implementations" ID="ID_327652489" CREATED="1493884238705" MODIFIED="1493884253919">
<node TEXT="Redis" ID="ID_1254944992" CREATED="1493884254480" MODIFIED="1493884256921">
<node TEXT="Hashtable: tasklist_workStationID" ID="ID_770453060" CREATED="1493884273500" MODIFIED="1493884368896"/>
<node TEXT="key: requestID" ID="ID_1493467728" CREATED="1493884370931" MODIFIED="1493884452253"/>
<node TEXT="value: serverID" ID="ID_1353059587" CREATED="1493884452559" MODIFIED="1493884457404"/>
</node>
<node TEXT="ZooKeeper: ?" ID="ID_1281363152" CREATED="1493884257293" MODIFIED="1493884263818"/>
<node TEXT="ETCD: ?" ID="ID_930819161" CREATED="1493884264076" MODIFIED="1493884270842"/>
</node>
</node>
<node TEXT="TakeOverTask(requestID, workstationID, oldServerID, newServerID) (jobDesc, error)" ID="ID_244771858" CREATED="1493883781297" MODIFIED="1493883905159">
<node TEXT="Redis:" ID="ID_1408799621" CREATED="1493886252440" MODIFIED="1493886263158">
<node TEXT="Solution 1" ID="ID_1587615298" CREATED="1493886263761" MODIFIED="1493886710814">
<node TEXT="Run MULTI" ID="ID_155993164" CREATED="1493886876730" MODIFIED="1493886897795"/>
<node TEXT="Run HGET tasklist_workStationID requestID" ID="ID_134563427" CREATED="1493886898180" MODIFIED="1493886915745"/>
<node TEXT="Run HDEL tasklist_workStationID requestID" ID="ID_1997216299" CREATED="1493886276779" MODIFIED="1493886298607"/>
<node TEXT="If the result is oldServerID and 1, Run HSETNX tasklist_workStationID requestID newServerID" ID="ID_156826902" CREATED="1493886299275" MODIFIED="1493886927062"/>
<node TEXT="else return error" ID="ID_1261678361" CREATED="1493886528995" MODIFIED="1493886532241"/>
</node>
<node TEXT="Solution 2" ID="ID_536473453" CREATED="1493886271878" MODIFIED="1493886274449">
<node TEXT="Lua script" ID="ID_762421668" CREATED="1493886429120" MODIFIED="1493886438743"/>
<node TEXT="val = HGET tasklist_workStationID requestID" ID="ID_661933234" CREATED="1493886439101" MODIFIED="1493886464730"/>
<node TEXT="if val == oldServerID: HSET tasklist_workStationID requestID newServerID" ID="ID_930232035" CREATED="1493886465071" MODIFIED="1493886515395"/>
<node TEXT="else return error" ID="ID_660040509" CREATED="1493886519997" MODIFIED="1493886524115"/>
</node>
</node>
<node TEXT="ZooKeeper:?" ID="ID_1774334160" CREATED="1493886547360" MODIFIED="1493886623915"/>
</node>
</node>
<node TEXT="Server" ID="ID_185646450" CREATED="1493882678591" MODIFIED="1493882681150">
<node TEXT="TCP ( including Unix )" ID="ID_1677416688" CREATED="1493883019622" MODIFIED="1493883032710">
<node TEXT="Serve( listener )" ID="ID_1994035189" CREATED="1493886977261" MODIFIED="1493886988351"/>
<node TEXT="GracefulStop( maxTimeout ) error" ID="ID_1244418234" CREATED="1493886989405" MODIFIED="1493886999778"/>
<node TEXT="Stop() error" ID="ID_1194631603" CREATED="1493887000299" MODIFIED="1493887003677"/>
<node TEXT="RegisterWorkflow( routes, iWorkflow )" ID="ID_1302664332" CREATED="1493887003902" MODIFIED="1493887021130"/>
<node TEXT="GetServerID() serverID" ID="ID_1702157544" CREATED="1493927082946" MODIFIED="1493927092212"/>
<node TEXT="SetCoordinator(iCoordinator)" ID="ID_724989922" CREATED="1493926335620" MODIFIED="1493926360909"/>
</node>
<node TEXT="UDP ( not support yet )" ID="ID_1867573112" CREATED="1493883033148" MODIFIED="1493883046858"/>
</node>
<node TEXT="Workflow" ID="ID_412784287" CREATED="1493882681986" MODIFIED="1493882690233">
<node TEXT="WorkStation" ID="ID_542349765" CREATED="1493926434233" MODIFIED="1493926438139">
<node TEXT="GetInputQueue() iQueue" ID="ID_289620666" CREATED="1493927316627" MODIFIED="1493927331444"/>
<node TEXT="SetWorker(iWorker, num)" ID="ID_1984436080" CREATED="1493927380633" MODIFIED="1493930443685"/>
</node>
<node TEXT="Workflow" ID="ID_10953647" CREATED="1493926438345" MODIFIED="1493926444475">
<node TEXT="HandleRequest( iRequest )" ID="ID_1224680348" CREATED="1493926445345" MODIFIED="1493926465235"/>
<node TEXT="RegisterWorkStation(iWorkStation)" ID="ID_364460348" CREATED="1493926474344" MODIFIED="1493926504443"/>
<node TEXT="LaunchWorkFlow()" ID="ID_1902460693" CREATED="1493927022691" MODIFIED="1493927078572"/>
</node>
<node TEXT="Worker" ID="ID_125115345" CREATED="1493930447837" MODIFIED="1493930454822">
<node TEXT="Do( iJob ) iJob, iError" ID="ID_514864914" CREATED="1493930455397" MODIFIED="1493932100957"/>
</node>
<node TEXT="Job" ID="ID_748974606" CREATED="1493932221303" MODIFIED="1493932225815"/>
</node>
<node TEXT="Request" ID="ID_1014334233" CREATED="1493918859835" MODIFIED="1493918868391">
<node TEXT="RequestID() requestID" ID="ID_1205652328" CREATED="1493925237897" MODIFIED="1493925253066"/>
<node TEXT="IsAsync() bool" ID="ID_1018808295" CREATED="1493926033308" MODIFIED="1493926040317"/>
<node TEXT="Client() iClient" ID="ID_147573663" CREATED="1493918868965" MODIFIED="1493924786296"/>
<node TEXT="IsClientAlive() bool" ID="ID_994365786" CREATED="1493926003869" MODIFIED="1493926013482"/>
<node TEXT="OriginRequest()(header []byte, body []byte)" ID="ID_1498413035" CREATED="1493924631690" MODIFIED="1493925181101"/>
<node TEXT="GetRequestResult() (header, body, error)" ID="ID_89172672" CREATED="1493925193530" MODIFIED="1493926106252"/>
<node TEXT="CurrentJob() iJob" ID="ID_1830734732" CREATED="1493926138953" MODIFIED="1493926166930"/>
</node>
<node TEXT="Error" ID="ID_671849169" CREATED="1493931987304" MODIFIED="1493931995297">
<node TEXT="ErrorCode() int" ID="ID_1333896989" CREATED="1493931995999" MODIFIED="1493932038298"/>
<node TEXT="ErrorMsg() string" ID="ID_1641355855" CREATED="1493932003487" MODIFIED="1493932014832"/>
</node>
<node TEXT="Client" ID="ID_587291771" CREATED="1493924531117" MODIFIED="1493924539929">
<node TEXT="io.Writer" ID="ID_1485039555" CREATED="1493924540925" MODIFIED="1493924583175"/>
<node TEXT="IsAlive() bool" ID="ID_1785887031" CREATED="1493924583836" MODIFIED="1493924630406"/>
</node>
</node>
<node TEXT="libs" POSITION="right" ID="ID_1446729204" CREATED="1493883091625" MODIFIED="1493883173328" HGAP_QUANTITY="25.249999664723884 pt" VSHIFT_QUANTITY="-40.49999879300598 pt">
<edge COLOR="#0000ff"/>
<node TEXT="Data Structures and Algorithms (dsa)" ID="ID_774848869" CREATED="1493883103794" MODIFIED="1493883119297"/>
<node TEXT="config" ID="ID_1751929998" CREATED="1493883119596" MODIFIED="1493883141556"/>
</node>
</node>
</map>