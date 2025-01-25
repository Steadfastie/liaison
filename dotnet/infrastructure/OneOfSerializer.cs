using MongoDB.Bson;
using MongoDB.Bson.IO;
using MongoDB.Bson.Serialization;
using MongoDB.Bson.Serialization.Serializers;
using OneOf;

namespace infrastructure;

public class OneOfTimeSpanDateTimeSerializer : SerializerBase<OneOf<TimeSpan, DateTime>>
{
    private const string Duration = "Duration";
    private const string ProcessedAt = "ProcessedAt";
    public override void Serialize(BsonSerializationContext context, BsonSerializationArgs args, OneOf<TimeSpan, DateTime> value)
    {
        if (value.IsT0) // TimeSpan
        {
            context.Writer.WriteStartDocument();
            context.Writer.WriteName(Duration);
            context.Writer.WriteInt64(value.AsT0.Ticks);
            context.Writer.WriteEndDocument();
        }
        else // DateTime
        {
            context.Writer.WriteStartDocument();
            context.Writer.WriteName(ProcessedAt);
            context.Writer.WriteDateTime(value.AsT1.Ticks);
            context.Writer.WriteEndDocument();
        }
    }

    public override OneOf<TimeSpan, DateTime> Deserialize(BsonDeserializationContext context, BsonDeserializationArgs args)
    {
        context.Reader.ReadStartDocument();
        string fieldName = context.Reader.ReadName();

        if (fieldName == Duration)
        {
            return new TimeSpan(context.Reader.ReadInt64());
        }
        else if (fieldName == ProcessedAt)
        {
            var timeValue = context.Reader.ReadDateTime();
            return new DateTime(timeValue);
        }
        else
        {
            throw new BsonSerializationException("Unexpected field name during deserialization");
        }
    }
}
