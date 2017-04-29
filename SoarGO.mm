<map version="1.0.1">
<!-- To view this file, download free mind mapping software FreeMind from http://freemind.sourceforge.net -->
<node CREATED="1493233878301" ID="ID_1905076007" MODIFIED="1493233885306" TEXT="SoarGO">
<node CREATED="1493233890149" ID="ID_1750066079" MODIFIED="1493233904055" POSITION="right" TEXT="libs">
<node CREATED="1493233904056" ID="ID_782298254" MODIFIED="1493233906668" TEXT="config">
<node CREATED="1493233906669" ID="ID_1078843836" MODIFIED="1493233929097" TEXT="Load config to struct"/>
<node CREATED="1493233929703" ID="ID_1483239267" MODIFIED="1493233941817" TEXT="Load from file"/>
<node CREATED="1493233942263" ID="ID_475169655" MODIFIED="1493233949521" TEXT="Load from env"/>
</node>
<node CREATED="1493233951739" ID="ID_352031776" MODIFIED="1493235964211" TEXT="Data Structure and Algorithms(dsa)">
<node CREATED="1493233974078" ID="ID_1308137064" MODIFIED="1493235644585" TEXT="lock free queue">
<node CREATED="1493233991172" ID="ID_620427042" MODIFIED="1493235714333" TEXT="Enqueue"/>
<node CREATED="1493235719284" ID="ID_1947385650" MODIFIED="1493235731732" TEXT="Dequeue">
<node CREATED="1493235732169" ID="ID_588594665" MODIFIED="1493235739204" TEXT="Block Until not empty"/>
</node>
</node>
</node>
<node CREATED="1493235967404" ID="ID_1789672059" MODIFIED="1493235970389" TEXT="sys">
<node CREATED="1493235970725" ID="ID_414329936" MODIFIED="1493236103611" TEXT="fdattr">
<node CREATED="1493235983522" ID="ID_1289685397" MODIFIED="1493236081105" TEXT="set fd not close when switches to new spawned process"/>
<node CREATED="1493236081931" ID="ID_213499463" MODIFIED="1493236089300" TEXT="set fd auto close when process switches"/>
</node>
<node CREATED="1493236104626" ID="ID_1030708490" MODIFIED="1493236111390" TEXT="inheritfd">
<node CREATED="1493236111391" ID="ID_1745510017" MODIFIED="1493236179811" TEXT="Register a fd to be inheritable. Supported fd type:">
<node CREATED="1493236179813" ID="ID_65170062" MODIFIED="1493236185606" TEXT="TCPListener"/>
<node CREATED="1493236185966" ID="ID_1287834096" MODIFIED="1493236195550" TEXT="UnixListener"/>
<node CREATED="1493236196026" ID="ID_79491584" MODIFIED="1493236199951" TEXT="UDPConn"/>
<node CREATED="1493236200386" ID="ID_383661946" MODIFIED="1493236204581" TEXT="File"/>
</node>
<node CREATED="1493236158936" ID="ID_633025094" MODIFIED="1493236227109" TEXT="Get the list of inheritable fds "/>
</node>
</node>
<node CREATED="1493437148439" ID="ID_927451007" MODIFIED="1493437150347" TEXT="app">
<node CREATED="1493449148476" ID="ID_1285652496" MODIFIED="1493449154688" TEXT="Servers">
<node CREATED="1493449241287" ID="ID_388942990" MODIFIED="1493449248720" TEXT="Listener">
<node CREATED="1493449250960" ID="ID_750903859" MODIFIED="1493449257561" TEXT="Generic Socket"/>
<node CREATED="1493449257829" ID="ID_598092456" MODIFIED="1493449261948" TEXT="Http(s)"/>
<node CREATED="1493449262268" ID="ID_730490110" MODIFIED="1493449269631" TEXT="GRPC"/>
</node>
<node CREATED="1493449271199" ID="ID_1410121612" MODIFIED="1493449277549" TEXT="Workflows">
<node CREATED="1493449277550" ID="ID_1648019791" MODIFIED="1493449348849" TEXT="Work Stations">
<node CREATED="1493449349402" ID="ID_850008326" MODIFIED="1493449353547" TEXT="Input Queue"/>
<node CREATED="1493449353818" ID="ID_813367397" MODIFIED="1493449358855" TEXT="N Workers"/>
<node CREATED="1493449359091" ID="ID_591950439" MODIFIED="1493449364223" TEXT="Worker Callback Function"/>
<node CREATED="1493449378969" ID="ID_1311967226" MODIFIED="1493449387674" TEXT="Dead Job Monitor"/>
</node>
<node CREATED="1493449300322" ID="ID_1228137643" MODIFIED="1493449634161" TEXT="Workflow manager">
<node CREATED="1493449581871" ID="ID_328615796" MODIFIED="1493449622774" TEXT="Connect the work stations"/>
<node CREATED="1493449624101" ID="ID_1288839893" MODIFIED="1493449697253" TEXT="Manage request types">
<node CREATED="1493449697254" ID="ID_422666714" MODIFIED="1493449710961" TEXT="Async ( Default)">
<node CREATED="1493449818748" ID="ID_597932128" MODIFIED="1493449866688" TEXT="return RequestReceived immediately">
<node CREATED="1493449866689" ID="ID_264095304" MODIFIED="1493449915726" TEXT="if https, return 202"/>
<node CREATED="1493449917064" ID="ID_384712743" MODIFIED="1493449967702" TEXT="Other types of server should also implement similar interface"/>
</node>
<node CREATED="1493449973185" ID="ID_1994870161" MODIFIED="1493450169217">
<richcontent TYPE="NODE"><html>
  <head>
    
  </head>
  <body>
    <p>
      When finish processing a request, if the client connection still exists, send back the result .&#160;
    </p>
  </body>
</html>
</richcontent>
</node>
</node>
<node CREATED="1493449699599" ID="ID_374853962" MODIFIED="1493449703343" TEXT="Sync">
<node CREATED="1493450205197" ID="ID_1187969299" MODIFIED="1493450353533" TEXT="Not return until finish processing the whole requests .  "/>
<node CREATED="1493450377283" ID="ID_423143927" MODIFIED="1493450394991" TEXT="Don&apos;t upgrade the progress of the request"/>
</node>
</node>
</node>
</node>
</node>
<node CREATED="1493449155020" ID="ID_1874842075" MODIFIED="1493449176834" TEXT="Pulse handler"/>
<node CREATED="1493449335130" ID="ID_919983111" MODIFIED="1493449341879" TEXT="Coordinator Instance"/>
<node CREATED="1493449166071" ID="ID_31164591" MODIFIED="1493449195880" TEXT="Inherit Manager"/>
</node>
<node CREATED="1493437156669" ID="ID_450517103" MODIFIED="1493437161947" TEXT="coordinator">
<node CREATED="1493437162987" ID="ID_671913984" MODIFIED="1493437178700" TEXT="Interface">
<node CREATED="1493437451637" ID="ID_1146786399" MODIFIED="1493440027522" TEXT="NewRequest() requestID ">
<node CREATED="1493440028410" ID="ID_1515105275" MODIFIED="1493440051577" TEXT="Opt 1: obtained by UUID algorithm"/>
<node CREATED="1493440052827" ID="ID_1115775352" MODIFIED="1493440078553" TEXT="Opt 2: obtained by underneath coordinator">
<node CREATED="1493440079331" ID="ID_1070386127" MODIFIED="1493440087826" TEXT="Redis: INCR requestID"/>
<node CREATED="1493440088567" ID="ID_1190128081" MODIFIED="1493440123693" TEXT="ZooKeeper: Create permanent sequential znode request_id_xxxxxxxxxx"/>
</node>
</node>
<node CREATED="1493438156796" ID="ID_558494582" MODIFIED="1493448815121" TEXT="UpdateRequestProgress( requestID, workstation, workflowID, inputMsg, msg) ">
<node CREATED="1493440589884" ID="ID_283988137" MODIFIED="1493440591244" TEXT="workflowID indicates the current owner of a task"/>
<node CREATED="1493440592619" ID="ID_424103264" MODIFIED="1493440622004" TEXT="A workstation relates to a specific tasks set"/>
</node>
<node CREATED="1493440271852" ID="ID_117436647" MODIFIED="1493449023147" TEXT="GetRequestProgress(requestID) workStation"/>
<node CREATED="1493440892444" ID="ID_1004844871" MODIFIED="1493440930964" TEXT="SaveRequestOutput(requestID, msg)"/>
<node CREATED="1493440931436" ID="ID_1590734327" MODIFIED="1493449041115" TEXT="GetRequestOutput(requestID) msg"/>
<node CREATED="1493440735628" ID="ID_218808568" MODIFIED="1493440749857" TEXT="GetTasksList(workstation)">
<node CREATED="1493440749857" ID="ID_332621967" MODIFIED="1493440759665" TEXT="Returns tasks list"/>
<node CREATED="1493440759920" ID="ID_1206880557" MODIFIED="1493440779228" TEXT="Each task contains owner info and launched time"/>
</node>
<node CREATED="1493440982124" ID="ID_891016372" MODIFIED="1493446356497" TEXT="TakeOverTask(requestID, workstation, oldWorkflowID, newWorkflowID) inputJson "/>
</node>
<node CREATED="1493437168384" ID="ID_1836724017" MODIFIED="1493437175513" TEXT="Implementation">
<node CREATED="1493437179640" ID="ID_1210067095" MODIFIED="1493437180779" TEXT="Redis"/>
<node CREATED="1493437181000" ID="ID_1543143883" MODIFIED="1493437183619" TEXT="ETCD"/>
<node CREATED="1493437183852" ID="ID_1447663047" MODIFIED="1493437186571" TEXT="ZooKeeper"/>
</node>
</node>
</node>
<node CREATED="1493236237806" HGAP="48" ID="ID_1181144637" MODIFIED="1493436962109" POSITION="left" TEXT="Architecture design" VSHIFT="-25">
<node CREATED="1493236252344" ID="ID_903117282" MODIFIED="1493236430414" TEXT="An App can have multiple servers"/>
<node CREATED="1493236433268" ID="ID_584625327" MODIFIED="1493236469221" TEXT="A Server consists of multiple workflows"/>
<node CREATED="1493236481629" ID="ID_1024356783" MODIFIED="1493237014274" TEXT="A workflow contains multiple workstations "/>
<node CREATED="1493236526611" ID="ID_1327820482" MODIFIED="1493237065062" TEXT="Each workstation focus on one simple mission and has multiple workers working on it  "/>
<node CREATED="1493236684555" ID="ID_288013443" MODIFIED="1493237190157" TEXT="Each worker will continuously do 5 things ">
<node CREATED="1493237190158" ID="ID_839145169" MODIFIED="1493237215253" TEXT="Grab a new task from the input queue "/>
<node CREATED="1493237216618" ID="ID_471365688" MODIFIED="1493237282696" TEXT="Mark the task as &quot;started&quot;"/>
<node CREATED="1493237250992" ID="ID_289009539" MODIFIED="1493237271107" TEXT="Process the task"/>
<node CREATED="1493237273976" ID="ID_713339557" MODIFIED="1493237301714" TEXT="Mark the task as &quot;finished&quot;"/>
<node CREATED="1493237302216" ID="ID_128053144" MODIFIED="1493237366629" TEXT="Put the result to output queue if it&apos;s not nil  "/>
</node>
</node>
<node CREATED="1493237512748" ID="ID_1760022409" MODIFIED="1493237569456" POSITION="right" TEXT="Possible technical issues">
<node CREATED="1493237571462" ID="ID_1404975407" MODIFIED="1493237626319" TEXT="Tasks management"/>
<node CREATED="1493276355569" ID="ID_126697518" MODIFIED="1493437040520" TEXT="Tasks Interrupt and Resume">
<node CREATED="1493277015123" ID="ID_882099883" MODIFIED="1493277067488" TEXT="Solution 1">
<node CREATED="1493277074164" ID="ID_470150676" MODIFIED="1493277095370" TEXT="Use ZooKeeper to manage locks for accessing a task"/>
<node CREATED="1493277095950" ID="ID_1846434586" MODIFIED="1493278726124" TEXT="A worker will create a ephemeral znode &amp; a permanent znode corresponding to the task "/>
<node CREATED="1493277116140" ID="ID_1089661252" MODIFIED="1493279030477" TEXT="The worker will delete both e&amp;p znode after finished executing the task"/>
<node CREATED="1493279031122" ID="ID_1067943557" MODIFIED="1493279425006" TEXT="If another node detect a p-znode of a task exist while the related e-znode is gone, it will take over and redo the task"/>
</node>
<node CREATED="1493277067952" ID="ID_234335942" MODIFIED="1493415749342" TEXT="Solution 2 (Coordinator Interface)">
<node CREATED="1493415750843" ID="ID_1165545863" MODIFIED="1493415876082" TEXT="Get tasks list"/>
<node CREATED="1493415876476" ID="ID_1850330539" MODIFIED="1493425498046" TEXT="Take over a task">
<node CREATED="1493428614390" ID="ID_775930788" MODIFIED="1493436397241" TEXT="Check if the task is being executed more than normal time"/>
<node CREATED="1493415970691" ID="ID_1476561244" MODIFIED="1493436415616" TEXT="If yes, then confirm  the current owner is dead">
<node CREATED="1493425512344" ID="ID_958960638" MODIFIED="1493425518321" TEXT="Try connect to it"/>
<node CREATED="1493425518630" ID="ID_150176741" MODIFIED="1493425534274" TEXT="Try PING-PONG pulse"/>
<node CREATED="1493425535083" ID="ID_49242648" MODIFIED="1493425569258" TEXT="Retry until reach MAX_RETRY_LIMIT"/>
<node CREATED="1493425570259" ID="ID_1267178839" MODIFIED="1493425588420" TEXT="Wait 10 seconds between each retry"/>
</node>
<node CREATED="1493425501097" ID="ID_1030862000" MODIFIED="1493425940346" TEXT="Try Compare-And-Swap the owner"/>
<node CREATED="1493427626039" ID="ID_644478129" MODIFIED="1493436434420" TEXT="Reset the last execution time"/>
</node>
<node CREATED="1493436482842" ID="ID_1668512737" MODIFIED="1493436489481" TEXT="Create a new task">
<node CREATED="1493436507293" ID="ID_536876635" MODIFIED="1493436569752" TEXT="Add task to the task list corresponding to a  set of workstation "/>
<node CREATED="1493436581101" ID="ID_1565791228" MODIFIED="1493436875450" TEXT="Save task owner and execution time info"/>
</node>
<node CREATED="1493436800252" ID="ID_203169236" MODIFIED="1493436808730" TEXT="Save request output"/>
<node CREATED="1493436809618" ID="ID_362544558" MODIFIED="1493436825419" TEXT="Update/Get request progress"/>
</node>
</node>
</node>
<node CREATED="1493237629445" ID="ID_1201833478" MODIFIED="1493268231861" POSITION="left" TEXT="Request handling">
<node CREATED="1493267362971" ID="ID_662187531" MODIFIED="1493267398431" TEXT="Each request will be associated with an unique ID"/>
<node CREATED="1493268185646" ID="ID_1573995568" MODIFIED="1493268270863" TEXT="A request involves multiple tasks"/>
<node CREATED="1493268350818" ID="ID_1549809120" MODIFIED="1493268407688" TEXT="Each task will be stored in a task list related to a workstation "/>
<node CREATED="1493268408791" ID="ID_353503776" MODIFIED="1493274918061" TEXT="A task is identified by node ID and request ID"/>
<node CREATED="1493274994414" ID="ID_477527690" MODIFIED="1493275022006" TEXT="A global table will be used to track the progress of a request"/>
</node>
<node CREATED="1493275295609" ID="ID_1179419688" MODIFIED="1493275302988" POSITION="left" TEXT="External requirement">
<node CREATED="1493275302989" ID="ID_344147365" MODIFIED="1493276227947" TEXT="System to store  requests progress and task status">
<node CREATED="1493276230096" ID="ID_132709629" MODIFIED="1493276232939" TEXT="Redis"/>
<node CREATED="1493276233193" ID="ID_339628104" MODIFIED="1493276282994" TEXT="ZooKeeper"/>
</node>
<node CREATED="1493276284169" ID="ID_1144144940" MODIFIED="1493276319349" TEXT="System to store task output">
<node CREATED="1493276319350" ID="ID_377069493" MODIFIED="1493276321636" TEXT="Redis"/>
<node CREATED="1493276327077" ID="ID_1336791476" MODIFIED="1493276329627" TEXT="Memcache"/>
<node CREATED="1493276329903" ID="ID_537016196" MODIFIED="1493276332163" TEXT="CouchDB"/>
<node CREATED="1493276332657" ID="ID_198994444" MODIFIED="1493276335211" TEXT="MongoDB"/>
</node>
</node>
<node CREATED="1493276454014" ID="ID_1070307027" MODIFIED="1493276458274" POSITION="right" TEXT="Health Check">
<node CREATED="1493276458275" ID="ID_713195230" MODIFIED="1493276476885" TEXT="Dynamic topology?"/>
<node CREATED="1493276486926" ID="ID_520956956" MODIFIED="1493276774342" TEXT="Pulse(PING-PONG)(SoarGo will listen at port 7627(?)to answer PING request)"/>
</node>
<node CREATED="1493436978403" ID="ID_965690397" MODIFIED="1493436994767" POSITION="left" TEXT="System features">
<node CREATED="1493436994768" ID="ID_66507630" MODIFIED="1493437067492" TEXT="0 downtime binary change"/>
<node CREATED="1493437069947" ID="ID_1794807003" MODIFIED="1493437081355" TEXT="Graceful shutdown"/>
<node CREATED="1493437081851" ID="ID_1481288140" MODIFIED="1493437109739" TEXT="Stateless"/>
<node CREATED="1493437110512" ID="ID_712516937" MODIFIED="1493437121371" TEXT="Scalable"/>
<node CREATED="1493437121696" ID="ID_675340905" MODIFIED="1493437123681" TEXT="Recoverable"/>
</node>
</node>
</map>
