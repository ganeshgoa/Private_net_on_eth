Install GO-Ethereum
sudo apt-get install software-properties-common
sudo add-apt-repository -y ppa:ethereum/ethereum
sudo apt-get update
sudo apt-get install ethereum
First launch bootnode

# Create node1:
geth account new --datadir node1
# Public address node1:
0x3E6671290ECbE283F9A8d46a39949bF95d41D0Eb

# Create node2:
geth account new --datadir node2
# Public address node2:
0x9B6deE802B98fe4A510F482c384F98A3651D44e4

# Create node3:
geth account new --datadir node3
# Public address node3:
0xb998BD1E1a87FE953c4ce90A3F76ECA53f78059e

# Create node4:
geth account new --datadir node4
# Public address node4:
0x83515F823dE176dd37143faFFD44aE58B5048517

# Create node5:
geth account new --datadir node5
# Public address node5:
0x7C39eb16aad6F7bf990873fd061304636C300006

# Bootnode address:
735137734c5e95e47f6b50c255b1137aea9cd06256d9eab3d1151f328eb7028263bc48441ef383c95944777378b119832f6f2f4697352026268311f4e47e6dfb@127.0.0.1:0?discport=30301

#Launch bootnode:
bootnode --nodekey=boot.key

# Launch node1 in new terminal
geth init --datadir node1 genesis.json
echo '75612' > node1/password.txt
geth --datadir node1 \
  --identity node1 \
  --syncmode full \
  --port 30311  \
  --ws \
  --ws.addr 0.0.0.0 \
  --ws.port 8546 \
  --ws.origins "*" \
  --http \
  --http.addr 0.0.0.0 \
  --http.port 8545 \
  --http.corsdomain "*" \
  --http.api shh,personal,db,eth,net,web3,txpool,miner,admin \
  --bootnodes enode://735137734c5e95e47f6b50c255b1137aea9cd06256d9eab3d1151f328eb7028263bc48441ef383c95944777378b119832f6f2f4697352026268311f4e47e6dfb@127.0.0.1:30301 \
  --networkid 777 \
  --unlock 0x3E6671290ECbE283F9A8d46a39949bF95d41D0Eb \
  --password node1/password.txt \
  --mine \
  --allow-insecure-unlock

# Launch node2:
geth init --datadir node2 genesis.json
echo '75612' > node2/password.txt
geth --datadir node2 \
  --identity node2 \
  --syncmode full \
  --port 30312  \
  --http.api shh,personal,db,eth,net,web3,txpool,miner,admin \
  --authrpc.port 8552 \
  --bootnodes enode://735137734c5e95e47f6b50c255b1137aea9cd06256d9eab3d1151f328eb7028263bc48441ef383c95944777378b119832f6f2f4697352026268311f4e47e6dfb@127.0.0.1:30301 \
  --networkid 777 \
  --unlock 0x9B6deE802B98fe4A510F482c384F98A3651D44e4 \
  --password node2/password.txt \
  --mine \
  --allow-insecure-unlock

# Launch node3:
geth init --datadir node3 genesis.json
echo '75612' > node3/password.txt
geth --datadir node3 \
  --identity node3 \
  --syncmode full \
  --port 30313  \
  --http.api shh,personal,db,eth,net,web3,txpool,miner,admin \
  --authrpc.port 8553 \
  --bootnodes enode://735137734c5e95e47f6b50c255b1137aea9cd06256d9eab3d1151f328eb7028263bc48441ef383c95944777378b119832f6f2f4697352026268311f4e47e6dfb@127.0.0.1:30301 \
  --networkid 777 \
  --unlock 0xb998BD1E1a87FE953c4ce90A3F76ECA53f78059e \
  --password node3/password.txt \
  --mine \
  --allow-insecure-unlock

# Launch node4:
geth init --datadir node4 genesis.json
echo '75612' > node4/password.txt
geth --datadir node4 \
  --identity node4 \
  --syncmode full \
  --port 30314  \
  --http.api shh,personal,db,eth,net,web3,txpool,miner,admin \
  --authrpc.port 8554 \
  --bootnodes enode://735137734c5e95e47f6b50c255b1137aea9cd06256d9eab3d1151f328eb7028263bc48441ef383c95944777378b119832f6f2f4697352026268311f4e47e6dfb@127.0.0.1:30301 \
  --networkid 777 \
  --unlock 0x83515F823dE176dd37143faFFD44aE58B5048517 \
  --password node4/password.txt \
  --mine \
  --allow-insecure-unlock

# Launch node4:
geth init --datadir node5 genesis.json
echo '75612' > node5/password.txt
geth --datadir node5 \
  --identity node5 \
  --syncmode full \
  --port 30315  \
  --http.api shh,personal,db,eth,net,web3,txpool,miner,admin \
  --authrpc.port 8555 \
  --bootnodes enode://735137734c5e95e47f6b50c255b1137aea9cd06256d9eab3d1151f328eb7028263bc48441ef383c95944777378b119832f6f2f4697352026268311f4e47e6dfb@127.0.0.1:30301 \
  --networkid 777 \
  --unlock 0x7C39eb16aad6F7bf990873fd061304636C300006 \
  --password node5/password.txt \
  --mine \
  --allow-insecure-unlock
