local function matchUsers(queueKey, pubSubChannel, minUsers, minScore, lobbyId, userId, userScore)
    local users = redis.call('ZRANGEBYSCORE', queueKey, minScore, '+inf', 'LIMIT', 0, minUsers)
    if #users < minUsers then

        for i, v in ipairs(users) do
            users[i] = tonumber(v)
        end


        table.insert(users, userId)
        redis.call('ZREM', queueKey, unpack(users))

        
        local lobby = {
            id = lobbyId,
            participants = users,
            created_at = UserScore
            state = 'started'
        }

        local lobbyJson = cjson.encode(lobby)
        redis.call('JSON.SET', 'lobby:' ..lobbyId, '.', lobbyJson)


        redis.call('PUBLISH', pubSubChannel, lobbyId .. ':' .. table.concat(users, ','))
        return {true, lobbyId, users}
    else
        redis.call('ZADD', queueKey, userScore, userId)
        return {false}
    end
    return {false}
end

local queueKey = KEYS[1]
local pubSubChannel = KEYS[2]
local minUsers = tonumber(ARGV[1])
local minScore = tonumber(ARGV[2])
local lobbyId = ARGV[3]
local userId = tonumber(ARGV[4])
local userScore = tonumber(ARGV[5])

return matchUsers(queueKey, pubSubChannel, minUsers, minScore, lobbyId, userId, userScore)
